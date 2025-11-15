package main

import (
	"container/heap"
	"fmt"
)

// Item representa um valor armazenado no heap mínimo.
type Item struct {
	Value int
}

// MinHeap implementa a interface heap.Interface para manter o menor elemento no topo.
type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// TopK retorna os K maiores valores de uma slice usando heap O(n log k).
func TopK(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return nil
	}

	h := &MinHeap{}
	heap.Init(h)

	for _, v := range nums {
		if h.Len() < k {
			// Preenche o heap até ter k elementos.
			heap.Push(h, Item{Value: v})
		} else if v > (*h)[0].Value {
			// Substitui o menor quando encontramos um número maior.
			heap.Pop(h)
			heap.Push(h, Item{Value: v})
		}
	}

	res := make([]int, h.Len())
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(Item).Value
	}
	return res
}

func main() {
	nums := []int{3, 10, 4, 7, 20, 15, 2, 9}
	fmt.Println("Top 3:", TopK(nums, 3))
}
