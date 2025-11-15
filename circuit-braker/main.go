package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrOpenCircuit   = errors.New("circuit breaker: open")
	ErrHalfOpenTrial = errors.New("circuit breaker: half-open trial in progress")
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu          sync.Mutex
	state       State
	failures    int
	threshold   int
	openUntil   time.Time
	openTimeout time.Duration
	trialInUse  bool
}

func NewCircuitBreaker(threshold int, openTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       Closed,
		threshold:   threshold,
		openTimeout: openTimeout,
	}
}

func (cb *CircuitBreaker) beforeCall() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()
	switch cb.state {
	case Closed:
		return nil
	case Open:
		if now.After(cb.openUntil) {
			cb.state = HalfOpen
			cb.trialInUse = false
			return nil
		}
		return ErrOpenCircuit
	case HalfOpen:
		if cb.trialInUse {
			return ErrHalfOpenTrial
		}
		cb.trialInUse = true
		return nil
	default:
		return nil
	}
}

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
			cb.failures = 0
		}
	case HalfOpen:
		if err != nil {
			cb.state = Open
			cb.openUntil = time.Now().Add(cb.openTimeout)
			fmt.Println("Circuito HALF-OPEN -> OPEN")
		} else {
			cb.state = Closed
			cb.failures = 0
			fmt.Println("Circuito HALF-OPEN -> CLOSED")
		}
		cb.trialInUse = false
	}
}

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
