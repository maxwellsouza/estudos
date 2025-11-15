package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data int
}

type Result struct {
	JobID int
	Out   int
	Err   error
}

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Simula trabalho
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			results <- Result{
				JobID: job.ID,
				Out:   job.Data * 2,
				Err:   nil,
			}
		}
	}
}

// NewWorkerPool cria um pool com N workers e retorna funções para enviar jobs e receber resultados.
func NewWorkerPool(ctx context.Context, numWorkers int) (chan<- Job, <-chan Result) {
	jobs := make(chan Job)                   // entrada
	results := make(chan Result, numWorkers) // bufferado pra reduzir bloqueio

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(ctx, i, jobs, results, &wg)
	}

	// fecha canal de resultados quando todos terminarem
	go func() {
		wg.Wait()
		close(results)
	}()

	return jobs, results
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs, results := NewWorkerPool(ctx, 4)

	// Envia jobs em outra goroutine
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case jobs <- Job{ID: i, Data: i + 1}:
			}
		}
		close(jobs)
	}()

	// Consome resultados
	for res := range results {
		if res.Err != nil {
			fmt.Printf("Job %d falhou: %v\n", res.JobID, res.Err)
			continue
		}
		fmt.Printf("Job %d => %d\n", res.JobID, res.Out)
	}
}
