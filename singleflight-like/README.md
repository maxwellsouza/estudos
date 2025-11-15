# Request Coalescing (tipo `singleflight`)

## O que é

Um mecanismo para **coalescer requisições**: múltiplas goroutines pedindo a mesma coisa ao mesmo tempo compartilham o resultado de uma única execução.

## Por que é útil

- Evitar _stampede_ em cache miss.
- Reduzir chamadas duplicadas a APIs externas ou banco.
- Padrão muito útil em serviços de alta concorrência.

## Como rodar

```bash
go run main.go
