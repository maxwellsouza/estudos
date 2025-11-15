package main

import (
	"fmt"
	"sync"
	"time"
)

// call armazena o resultado compartilhado e sincroniza leitores concorrentes.
type call struct {
	wg  sync.WaitGroup
	val any
	err error
}

// Group mapeia chaves para chamadas em andamento, evitando duplicação de trabalho.
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do garante que chamadas concorrentes com a mesma chave compartilhem o mesmo resultado.
func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		// Já existe uma operação em andamento; aguarda o resultado.
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// Executa a função lentamente (ou remotamente) fora do lock.
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

func main() {
	var g Group

	fn := func() (any, error) {
		fmt.Println("Executando operação pesada...")
		time.Sleep(1 * time.Second)
		return "resultado", nil
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val, err := g.Do("chave-unica", fn)
			fmt.Printf("Goroutine %d recebeu: %v (err=%v)\n", id, val, err)
		}(i)
	}

	wg.Wait()
}
