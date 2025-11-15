# Worker Pool com Context e Backpressure

## O que é

Um pool de workers concorrentes que processa jobs em paralelo usando goroutines e canais. Usa `context.Context` para cancelamento e timeout.

## Por que é útil

- Processamento paralelo de filas (e-mails, PDFs, eventos, webhooks).
- Controle de concorrência para não sobrecarregar banco ou APIs.
- Cancelamento elegante em shutdown de serviço.

## Como rodar

```bash
go run main.go
