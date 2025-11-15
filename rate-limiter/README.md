# Rate Limiter (Token Bucket)

## O que é

Um limitador de taxa simples baseado em _token bucket_. Controla quantas operações podem ocorrer por segundo e suporta "burst".

## Por que é útil

- Evitar estourar limites de APIs externas.
- Proteger seu próprio backend de overload.
- Limitar requisições de usuários ou jobs.

## Como rodar

```bash
go run main.go
