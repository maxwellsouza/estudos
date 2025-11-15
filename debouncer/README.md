# Debouncer de Funções

## O que é

Um **debouncer** garante que uma função só será executada após um período de inatividade — ou seja, após o último disparo passar X milissegundos sem novos eventos.

## Por que é útil

- Evitar chamadas em excesso (ex: busca enquanto o usuário digita).
- Reduzir carga em APIs ou banco ao reagir a eventos frequentes.
- Muito usado em UIs e sistemas reativos.

## Como rodar

```bash
go run main.go
