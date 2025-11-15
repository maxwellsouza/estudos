package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ErrMaxRetries indica que todas as tentativas foram esgotadas.
var ErrMaxRetries = errors.New("max retries reached")

type Operation func(ctx context.Context) error

// Retry executa uma operação até que ela retorne sucesso, respeitando o
// número máximo de tentativas e aplicando exponencial backoff com jitter.
func Retry(ctx context.Context, op Operation, maxRetries int, baseDelay time.Duration) error {
	var attempt int
	for {
		err := op(ctx)
		if err == nil {
			return nil
		}

		attempt++
		if attempt > maxRetries {
			return fmt.Errorf("%w: last error: %v", ErrMaxRetries, err)
		}

		// Exponential backoff com jitter
		backoff := float64(baseDelay) * math.Pow(2, float64(attempt-1))
		jitter := rand.Float64() * float64(baseDelay)
		sleep := time.Duration(backoff + jitter)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(sleep):
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tries := 0
	err := Retry(ctx, func(ctx context.Context) error {
		tries++
		fmt.Println("Tentativa", tries)
		if tries < 3 {
			return errors.New("falha temporária")
		}
		fmt.Println("Sucesso!")
		return nil
	}, 5, 200*time.Millisecond)
	if err != nil {
		fmt.Println("Erro final:", err)
	}
}
