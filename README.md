
# 🧠 Estudos em Go — Algoritmos Essenciais

Este repositório reúne uma coleção de **algoritmos, padrões de concorrência e estruturas avançadas em Go**, organizados por dificuldade e aplicabilidade no mundo real.

O objetivo é servir como:

* **Playground de aprendizado contínuo**
* **Base de referência rápida** para soluções idiomáticas em Go
* **Portfólio técnico** com exemplos sólidos e bem estruturados

Cada algoritmo está em sua própria pasta, contendo:

✔ `main.go` — implementação + exemplo
✔ `README.md` — explicações, casos de uso e instruções de execução

---

## 📂 Lista de Algoritmos

| #  | Algoritmo / Padrão                     | Dificuldade    | Tema Central              | Descrição                                                    |
| -- | -------------------------------------- | -------------- | ------------------------- | ------------------------------------------------------------ |
| 01 | LRU Cache Genérico Thread-Safe         | Médio          | Estruturas e concorrência | Cache com remoção do menos usado via map + lista ligada      |
| 02 | Worker Pool com Context e Backpressure | Médio          | Concorrência              | Execução paralela segura, cancelável e controlada            |
| 03 | Rate Limiter (Token Bucket)            | Médio          | Governança/Resiliência    | Controle de taxa por segundo e suporte a burst               |
| 04 | Retry com Exponential Backoff          | Médio          | Resiliência               | Re-tentativas com aumento exponencial e jitter               |
| 05 | Circuit Breaker Simples                | Avançado       | Arquitetura resiliente    | Padrão clássico para microserviços e APIs instáveis          |
| 06 | Pipeline Concorrente (Fan-out/Fan-in)  | Avançado       | Concorrência              | Encadeamento concorrente de estágios de processamento        |
| 07 | Top-K com Heap Mínimo                  | Médio          | Estruturas                | Seleção eficiente dos maiores elementos                      |
| 08 | Dijkstra (Menor Caminho)               | Avançado       | Grafos                    | Cálculo de caminhos mínimos com heap de prioridade           |
| 09 | Debouncer de Funções                   | Médio          | Performance/UI/API        | Executa função só após período de inatividade                |
| 10 | Request Coalescing (tipo singleflight) | Muito avançado | Concorrência              | Unifica múltiplas requisições simultâneas em uma só execução |

---

## ▶️ Como executar qualquer algoritmo

Entre na pasta desejada:

```bash
cd lru-cache
```

E execute:

```bash
go run main.go
```

Todos os exemplos foram feitos com:

* Go 1.20+
* Apenas biblioteca padrão (`stdlib`)
* Código limpo, idiomático e fácil de adaptar

---

## 🎯 Objetivo

Este repositório foi projetado para:

* Dominar padrões de concorrência e paralelismo em Go
* Entender estruturas amplamente usadas no mercado
* Aprender técnicas avançadas como circuit breaker, retry e pipelines
* Construir um portfólio sólido com exemplos aplicáveis

---