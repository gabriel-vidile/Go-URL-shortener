# Go URL Shortener

**go-url-shortner** é uma aplicação simples escrita em Go que permite encurtar URLs longas e redirecionar para elas usando URLs curtas. Ela utiliza SQLite como banco de dados para persistência e inclui um sistema de cache para melhorar o desempenho.

---

## **Índice**
1. [Funcionalidades](#funcionalidades)
2. [Requisitos](#requisitos)
3. [Configuração e Execução](#configuração-e-execução)
   - [Clonar o Repositório](#1-clonar-o-repositório)
   - [Configurar o Banco de Dados](#2-configurar-o-banco-de-dados)
   - [Executar Localmente](#3-executar-localmente)
   - [Executar com Docker](#4-executar-com-docker)
4. [Uso da API](#uso-da-api)
   - [Encurtar uma URL](#encurtar-uma-url)
   - [Redirecionar pela URL Curta](#redirecionar-pela-url-curta)
5. [Testes](#testes)
6. [Melhorias Futuras](#melhorias-futuras)

---

## **Funcionalidades**

- Encurtar URLs longas.
- Redirecionar para URLs originais usando URLs curtas.
- Persistência das URLs no banco de dados SQLite.
- Cache em memória para respostas rápidas.
- Sistema pronto para produção usando Docker.

---

## **Requisitos**

Antes de começar, certifique-se de ter instalado:

- [Go](https://go.dev/dl/) (v1.20 ou superior)
- [SQLite](https://www.sqlite.org/download.html) (opcional, apenas para interações diretas com o banco)
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) (para execução com container)

---

## **Configuração e Execução**

### **1. Clonar o Repositório**

Clone o repositório para sua máquina local:
```bash
git clone https://github.com/gabriel-vidile/go-url-shortner.git
cd go-url-shortner
