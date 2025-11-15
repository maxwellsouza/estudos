package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	// Fan-out
	w1 := worker(in)
	w2 := worker(in)

	// Fan-in
	for res := range merge(w1, w2) {
		fmt.Println("resultado:", res)
	}
}
