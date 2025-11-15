# Top-K com Heap (Priority Queue)

## O que é

Uma implementação de **Top K elementos** usando `container/heap` para manter um heap mínimo de tamanho `k`.

## Por que é útil

- Ranking de itens (maior score, maiores vendas, etc.).
- Quando o conjunto é grande, mas você só precisa dos poucos maiores.
- O(n log k) em vez de O(n log n).

## Como rodar

```bash
go run main.go
