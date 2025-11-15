package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ErrOpenCircuit e ErrHalfOpenTrial representam estados em que a chamada deve ser
// bloqueada pelo breaker.
var (
	ErrOpenCircuit   = errors.New("circuit breaker: open")
	ErrHalfOpenTrial = errors.New("circuit breaker: half-open trial in progress")
)

// State modela os estados possíveis do circuit breaker.
type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

// CircuitBreaker encapsula a lógica de transição de estados e contabilização
// de falhas. Os campos são protegidos por um mutex simples para suportar
// acesso concorrente.
type CircuitBreaker struct {
	mu          sync.Mutex
	state       State
	failures    int
	threshold   int
	openUntil   time.Time
	openTimeout time.Duration
	trialInUse  bool
}

// NewCircuitBreaker cria uma instância com o número máximo de falhas toleradas
// e o tempo de espera antes de permitir uma nova tentativa.
func NewCircuitBreaker(threshold int, openTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       Closed,
		threshold:   threshold,
		openTimeout: openTimeout,
	}
}

// beforeCall decide se uma chamada pode prosseguir de acordo com o estado
// atual do breaker. Quando o estado é OPEN ele verifica se o tempo de espera
// expirou para migrar para HALF-OPEN.
func (cb *CircuitBreaker) beforeCall() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()
	switch cb.state {
	case Closed:
		// Circuito fechado: chamadas liberadas normalmente.
		return nil
	case Open:
		if now.After(cb.openUntil) {
			// Tempo de espera acabou, permite uma chamada de teste.
			cb.state = HalfOpen
			cb.trialInUse = false
			return nil
		}
		return ErrOpenCircuit
	case HalfOpen:
		if cb.trialInUse {
			// Já existe uma chamada de teste em andamento.
			return ErrHalfOpenTrial
		}
		cb.trialInUse = true
		return nil
	default:
		return nil
	}
}

// afterCall atualiza o estado do breaker conforme o resultado da operação.
// Falhas sucessivas fecham o circuito; um sucesso em HALF-OPEN retorna para
// CLOSED.
func (cb *CircuitBreaker) afterCall(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case Closed:
		if err != nil {
			cb.failures++
			if cb.failures >= cb.threshold {
				cb.state = Open
				cb.openUntil = time.Now().Add(cb.openTimeout)
				fmt.Println("Circuito -> OPEN")
			}
		} else {
			// Sucesso limpa o contador de falhas.
			cb.failures = 0
		}
	case HalfOpen:
		if err != nil {
			// Falha durante tentativa promove retorno para OPEN.
			cb.state = Open
			cb.openUntil = time.Now().Add(cb.openTimeout)
			fmt.Println("Circuito HALF-OPEN -> OPEN")
		} else {
			// Sucesso restabelece o circuito para CLOSED.
			cb.state = Closed
			cb.failures = 0
			fmt.Println("Circuito HALF-OPEN -> CLOSED")
		}
		cb.trialInUse = false
	}
}

// Execute envolve uma função com a proteção do circuit breaker.
func (cb *CircuitBreaker) Execute(fn func() error) error {
	if err := cb.beforeCall(); err != nil {
		return err
	}
	err := fn()
	cb.afterCall(err)
	return err
}

// Exemplo de uso
func main() {
	cb := NewCircuitBreaker(3, 2*time.Second)

	fail := true
	for i := 0; i < 15; i++ {
		err := cb.Execute(func() error {
			if fail {
				// Simula falha do serviço protegido.
				return errors.New("falha no serviço")
			}
			return nil
		})

		fmt.Printf("[%d] err = %v, state = %v\n", i, err, cb.state)

		if i == 6 {
			fmt.Println("Simulando serviço recuperado...")
			fail = false
		}
		time.Sleep(300 * time.Millisecond)
	}
}
