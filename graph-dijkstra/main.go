package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge representa uma aresta com destino e peso.
type Edge struct {
	To     int
	Weight float64
}

// Graph mapeia um identificador de nó para suas arestas adjacentes.
type Graph map[int][]Edge

// NodeDist guarda o nó atual e a distância mínima conhecida para ele.
type NodeDist struct {
	Node int
	Dist float64
}

// MinPQ implementa uma fila de prioridade mínima usando o pacote heap.
type MinPQ []NodeDist

func (pq MinPQ) Len() int           { return len(pq) }
func (pq MinPQ) Less(i, j int) bool { return pq[i].Dist < pq[j].Dist }
func (pq MinPQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MinPQ) Push(x any)        { *pq = append(*pq, x.(NodeDist)) }
func (pq *MinPQ) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// Dijkstra calcula as menores distâncias a partir de um nó inicial em um grafo.
func Dijkstra(g Graph, start int) map[int]float64 {
	dist := make(map[int]float64)
	for node := range g {
		dist[node] = math.Inf(1)
	}
	dist[start] = 0

	pq := &MinPQ{}
	heap.Init(pq)
	heap.Push(pq, NodeDist{Node: start, Dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(NodeDist)
		if cur.Dist > dist[cur.Node] {
			// Um caminho melhor já foi encontrado; ignora este.
			continue
		}
		for _, e := range g[cur.Node] {
			nd := cur.Dist + e.Weight
			if nd < dist[e.To] {
				dist[e.To] = nd
				heap.Push(pq, NodeDist{Node: e.To, Dist: nd})
			}
		}
	}
	return dist
}

func main() {
	g := Graph{
		1: {{To: 2, Weight: 2}, {To: 3, Weight: 5}},
		2: {{To: 3, Weight: 1}, {To: 4, Weight: 2}},
		3: {{To: 4, Weight: 3}},
		4: {},
	}

	dist := Dijkstra(g, 1)
	fmt.Println("Distâncias a partir do nó 1:", dist)
}
