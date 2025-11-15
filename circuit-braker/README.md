# Circuit Breaker Simples

## O que é

Um **circuit breaker** controla chamadas a um serviço externo. Após várias falhas, ele "abre" o circuito e passa a recusar chamadas até um tempo de resfriamento.

## Por que é útil

- Evita sobrecarregar serviços já em falha.
- Acelera o *fail fast*.
- É um padrão clássico de resiliência em microserviços.

## Como rodar

```bash
go run main.go
