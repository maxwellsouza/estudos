package main

import (
	"fmt"
	"sync"
	"time"
)

type Debouncer struct {
	mu     sync.Mutex
	delay  time.Duration
	timer  *time.Timer
	action func()
}

func NewDebouncer(delay time.Duration, action func()) *Debouncer {
	return &Debouncer{
		delay:  delay,
		action: action,
	}
}

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

	for i := 0; i < 5; i++ {
		debouncedPrint.Call()
		time.Sleep(150 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)
}
