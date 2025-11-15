package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generator cria um canal que emite os números fornecidos sequencialmente.
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// worker consome valores de um canal de entrada e devolve o quadrado de cada
// elemento após um pequeno atraso aleatório para simular trabalho pesado.
func worker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			out <- n * n
		}
	}()
	return out
}

// merge realiza o fan-in combinando múltiplos canais em um único fluxo de saída.
func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, ch := range chs {
			for v := range ch {
				out <- v
			}
		}
	}()
	return out
}

func main() {
	in := generator(1, 2, 3, 4, 5, 6, 7, 8)

	// Fan-out: distribui o trabalho entre duas goroutines.
	w1 := worker(in)
	w2 := worker(in)

	// Fan-in: agrega as respostas dos workers no mesmo canal.
	for res := range merge(w1, w2) {
		fmt.Println("resultado:", res)
	}
}
