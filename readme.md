
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
git clone https://github.com/seu-usuario/go-url-shortner.git
cd go-url-shortner
```

---

### **2. Configurar o Banco de Dados**

O banco de dados SQLite será criado automaticamente na primeira execução. No entanto, você pode criar manualmente a estrutura do banco com o arquivo de migração.

#### Executar Migrações (Opcional)
Certifique-se de ter o SQLite instalado e execute:
```bash
sqlite3 url_shortener.db < database/migrations.sql
```

---

### **3. Executar Localmente**

1. Instale as dependências do projeto:
   ```bash
   go mod tidy
   ```

2. Execute o servidor:
   ```bash
   go run main.go
   ```

3. O servidor estará disponível em:
   ```
   http://localhost:8080
   ```

---

### **4. Executar com Docker**

#### Build da Imagem Docker
Crie a imagem Docker:
```bash
docker build -t go-url-shortner .
```

#### Executar o Container
Execute o container:
```bash
docker run -p 8080:8080 go-url-shortner
```

Agora o servidor estará disponível em:
```
http://localhost:8080
```

---

## **Uso da API**

A aplicação expõe dois endpoints principais:

### **1. Encurtar uma URL**

- **Endpoint**: `POST /shorten`
- **Body**: Envie um JSON com a URL a ser encurtada.
- **Exemplo de Requisição**:
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:8080/shorten
  ```
- **Resposta**:
  ```json
  {
    "short_url": "http://localhost:8080/abcd1234"
  }
  ```

### **2. Redirecionar pela URL Curta**

- **Endpoint**: `GET /:shortID`
- **Exemplo**: Acesse `http://localhost:8080/abcd1234` no navegador.
- **Comportamento**: Redireciona automaticamente para a URL original (`https://example.com`).

---

## **Testes**

Para rodar os testes unitários:

1. Execute os testes com:
   ```bash
   go test ./tests/...
   ```

2. Verifique os resultados para garantir que todos os componentes funcionem corretamente.

---

## **Melhorias Futuras**

- **Autenticação**: Adicionar um sistema de autenticação para gerenciar URLs encurtadas.
- **Expiração de URLs**: Implementar URLs temporárias que expiram após um período definido.
- **Métricas**: Coletar dados como número de cliques em cada URL.
- **Integração com Redis**: Substituir o cache em memória por uma solução distribuída.
- **Frontend**: Criar uma interface gráfica para o sistema.

---

## **Contribuições**

Contribuições são bem-vindas! Abra uma issue ou envie um pull request.

---

## **Licença**

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.
