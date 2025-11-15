package main

import (
	"fmt"
	"sync"
	"time"
)

// Debouncer acumula chamadas rápidas e garante que a ação seja executada
// apenas após um período de silêncio.
type Debouncer struct {
	mu     sync.Mutex
	delay  time.Duration
	timer  *time.Timer
	action func()
}

// NewDebouncer configura a estrutura com o intervalo desejado e a função que
// será chamada após o repouso.
func NewDebouncer(delay time.Duration, action func()) *Debouncer {
	return &Debouncer{
		delay:  delay,
		action: action,
	}
}

// Call reinicia o temporizador e agenda a execução da ação caso nenhuma nova
// chamada ocorra dentro do intervalo configurado.
func (d *Debouncer) Call() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.delay, func() {
		d.mu.Lock()
		defer d.mu.Unlock()
		d.action()
	})
}

func main() {
	debouncedPrint := NewDebouncer(500*time.Millisecond, func() {
		fmt.Println("Executado após 500ms sem novas chamadas")
	})

	// Dispara múltiplas chamadas em sequência, apenas a última executará.
	for i := 0; i < 5; i++ {
		debouncedPrint.Call()
		time.Sleep(150 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)
}
