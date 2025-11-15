<!-- file: 01-lru-cache/README.md -->
# LRU Cache Genérico (thread-safe)

## O que é

Uma implementação de cache LRU (Least Recently Used) genérico em Go, usando `map` + `list.List` e `sync.Mutex` para segurança entre goroutines.

## Por que é útil

- Cache de resultados de chamadas caras (banco, API externa, cálculos).
- Evita usar memória infinita: o item menos usado é descartado primeiro.
- Padrão muito comum em backends e sistemas de alta performance.

## Como usar

```bash
go run main.go
