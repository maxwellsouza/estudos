package main

import (
	"fmt"
	"time"
)

// TokenBucket simples com taxa fixa.
type TokenBucket struct {
	capacity int
	tokens   int
	ticker   *time.Ticker
	stop     chan struct{}
}

// NewTokenBucket inicia o reabastecimento periódico do bucket e devolve a
// estrutura pronta para uso.
func NewTokenBucket(ratePerSec, capacity int) *TokenBucket {
	tb := &TokenBucket{
		capacity: capacity,
		tokens:   capacity,
		ticker:   time.NewTicker(time.Second / time.Duration(ratePerSec)),
		stop:     make(chan struct{}),
	}

	go func() {
		for {
			select {
			case <-tb.ticker.C:
				if tb.tokens < tb.capacity {
					tb.tokens++
				}
			case <-tb.stop:
				return
			}
		}
	}()

	return tb
}

// Allow devolve se há tokens disponíveis para consumir.
func (tb *TokenBucket) Allow() bool {
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// Stop encerra a goroutine interna e libera recursos.
func (tb *TokenBucket) Stop() {
	close(tb.stop)
	tb.ticker.Stop()
}

func main() {
	bucket := NewTokenBucket(5, 10) // 5 req/s máx, 10 de burst
	defer bucket.Stop()

	for i := 0; i < 20; i++ {
		if bucket.Allow() {
			fmt.Println("Request", i, "permitido")
		} else {
			fmt.Println("Request", i, "bloqueado")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
