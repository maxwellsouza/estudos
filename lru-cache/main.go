package main

import (
	"container/list"
	"fmt"
	"sync"
)

// Cache é um LRU genérico, seguro para uso com múltiplas goroutines.
type Cache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	items    map[K]*list.Element
	list     *list.List // frente = mais recente, fundo = menos recente
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("capacity must be > 0")
	}
	return &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
		list:     list.New(),
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zero V

	elem, ok := c.items[key]
	if !ok {
		return zero, false
	}

	c.list.MoveToFront(elem)
	return elem.Value.(entry[K, V]).value, true
}

func (c *Cache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Atualiza se já existe
	if elem, ok := c.items[key]; ok {
		elem.Value = entry[K, V]{key, value}
		c.list.MoveToFront(elem)
		return
	}

	// Insere novo
	e := entry[K, V]{key, value}
	elem := c.list.PushFront(e)
	c.items[key] = elem

	// Remove o menos recente se passar da capacidade
	if c.list.Len() > c.capacity {
		last := c.list.Back()
		if last != nil {
			c.list.Remove(last)
			kv := last.Value.(entry[K, V])
			delete(c.items, kv.key)
		}
	}
}

func main() {
	cache := NewCache[string, int](3)

	cache.Put("a", 1)
	cache.Put("b", 2)

	if v, ok := cache.Get("a"); ok {
		fmt.Println("a =", v) // 1
	}

	// "c" entra, "b" ou "a" será removido (quem for menos usado)
	cache.Put("c", 3)

	if _, ok := cache.Get("b"); !ok {
		fmt.Println("b foi removido do cache LRU")
	}
}
