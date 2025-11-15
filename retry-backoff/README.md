# Retry com Exponential Backoff e Context

## O que é

Um helper de **re-tentativa** que executa uma operação repetidamente em caso de erro, usando *exponential backoff* com jitter e `context.Context` para cancelamento.

## Por que é útil

- Requisições HTTP instáveis.
- Operações de I/O que falham de forma intermitente.
- Reduz carga em sistemas já degradados.

## Como rodar

```bash
go run main.go
