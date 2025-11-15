# Pipeline Concorrente (Fan-out/Fan-in)

## O que é

Um exemplo de pipeline concorrente: um gerador envia dados, múltiplos workers processam em paralelo (fan-out) e depois os resultados são combinados (fan-in).

## Por que é útil

- Processamento de lotes de dados.
- Transformações encadeadas (ETL, streams).
- Padrão muito idiomático em Go usando canais.

## Como rodar

```bash
go run main.go
