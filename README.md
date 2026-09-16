# DevBook

O **DevBook** é um ecossistema de aplicações desenvolvido em **Go** que simula uma rede social voltada para desenvolvedores.

O projeto foi estruturado em duas aplicações principais:

- **API**: responsável pelas regras de negócio, autenticação, usuários, publicações e persistência dos dados.
- **WebApp**: responsável pela interface web, interação com o usuário, gerenciamento da sessão e comunicação com a API.

A proposta do projeto é demonstrar, de forma prática, como diferentes camadas de uma aplicação podem ser separadas e integradas através de uma API REST.

## 🧩 Arquitetura

```text
                         ┌──────────────────┐
                         │      Usuário     │
                         └────────┬─────────┘
                                  │
                                  ▼
                     ┌────────────────────────┐
                     │       DevBook          │
                     │        WebApp           │
                     │                        │
                     │ Go + HTML + JavaScript │
                     └────────────┬───────────┘
                                  │
                             HTTP / JSON
                                  │
                                  ▼
                     ┌────────────────────────┐
                     │       DevBook API      │
                     │                        │
                     │ Go + REST + JWT        │
                     └────────────┬───────────┘
                                  │
                              SQL / MySQL
                                  │
                                  ▼
                     ┌────────────────────────┐
                     │        MySQL           │
                     └────────────────────────┘
```

A separação permite que a camada de apresentação e a camada de backend evoluam de forma independente.

## 📦 Estrutura do projeto

```text
devbook/
├── api/
│   ├── main.go
│   ├── go.mod
│   ├── sql/
│   └── src/
│
├── webapp/
│   ├── main.go
│   ├── go.mod
│   ├── assets/
│   ├── views/
│   └── src/
│
└── README.md
```

### API

A pasta `api/` contém o backend da aplicação.

Entre suas responsabilidades estão:

- Cadastro e gerenciamento de usuários
- Login e autenticação
- Geração e validação de tokens JWT
- Hash de senhas com bcrypt
- Sistema de seguidores
- Criação e gerenciamento de publicações
- Curtidas
- Feed de publicações
- Persistência dos dados no MySQL
- Validação e tratamento de erros
- Exposição dos endpoints REST

### WebApp

A pasta `webapp/` contém a aplicação web que interage diretamente com o usuário.

Suas responsabilidades incluem:

- Renderização das páginas HTML
- Tela de login
- Cadastro de usuários
- Página principal
- Criação de publicações
- Curtidas e remoção de curtidas
- Gerenciamento da sessão
- Armazenamento seguro do token em cookie
- Comunicação HTTP com a API
- Interações assíncronas utilizando JavaScript e jQuery

## 🔄 Como as aplicações trabalham juntas

O fluxo básico do DevBook acontece da seguinte forma:

1. O usuário acessa o **WebApp**.
2. O WebApp apresenta as páginas utilizando templates HTML.
3. O usuário realiza login ou cadastro.
4. O WebApp envia a solicitação para a **DevBook API**.
5. A API valida os dados e executa as regras de negócio.
6. Em operações de autenticação, a API retorna um token JWT.
7. O WebApp armazena o token em um cookie protegido.
8. Nas operações autenticadas, o WebApp recupera o token.
9. O token é enviado para a API através do header `Authorization`.
10. A API consulta ou altera os dados no MySQL.
11. O resultado retorna para o WebApp.
12. O WebApp apresenta o resultado ao usuário.

## 🔐 Autenticação

A autenticação é centralizada na API utilizando **JWT**.

O WebApp funciona como cliente da API e mantém o token de autenticação em um cookie protegido utilizando `gorilla/securecookie`.

O fluxo pode ser representado assim:

```text
Login
  │
  ▼
WebApp
  │
  │ POST /login
  ▼
API
  │
  │ valida usuário
  │
  │ gera JWT
  ▼
WebApp
  │
  │ salva token
  ▼
Cookie seguro
```

Nas requisições protegidas:

```text
WebApp
  │
  │ Authorization: Bearer JWT
  ▼
API
  │
  │ valida token
  ▼
Regra de negócio
```

## 🛠️ Principais tecnologias

### Backend

- Go 1.25.2
- Gorilla Mux
- MySQL
- JWT
- bcrypt
- GoDotEnv
- Gorilla SecureCookie

### Web

- Go `html/template`
- HTML5
- CSS3
- JavaScript
- jQuery
- Bootstrap
- Font Awesome

### Comunicação

- HTTP
- JSON
- API REST

## 🎯 Objetivo do projeto

Mais do que uma rede social fictícia, o DevBook foi estruturado como um projeto prático para explorar conceitos fundamentais de desenvolvimento de software.

Entre os principais conceitos demonstrados estão:

- Desenvolvimento de APIs REST
- Arquitetura cliente-servidor
- Separação de responsabilidades
- Organização de aplicações em Go
- Controllers e middlewares
- Roteamento HTTP
- Consumo de APIs
- Autenticação e autorização
- JWT
- Cookies seguros
- Hash de senhas
- Persistência com MySQL
- Operações CRUD
- Relacionamentos entre usuários
- Templates server-side
- AJAX
- Variáveis de ambiente
- Comunicação entre aplicações independentes

## 🚀 Executando o ecossistema

O DevBook é composto por duas aplicações que devem estar configuradas para se comunicarem.

### 1. API

Acesse:

```bash
cd api
```

Configure o ambiente e execute:

```bash
go run .
```

Por padrão, a API utiliza a porta:

```text
5000
```

### 2. WebApp

Em outro terminal:

```bash
cd webapp
```

Configure o arquivo `.env` com a URL e porta da API e execute:

```bash
go run .
```

Por padrão, o WebApp utiliza a porta:

```text
4000
```

Depois, acesse:

```text
http://localhost:4000
```

## 🔗 Comunicação local

Uma configuração típica de desenvolvimento é:

```text
WebApp
http://localhost:4000
        │
        │ HTTP
        ▼
API
http://localhost:5000
        │
        │ SQL
        ▼
MySQL
localhost:3306
```

O endereço da API utilizado pelo WebApp é configurado através das variáveis de ambiente.

## 📚 Documentação específica

Cada aplicação possui sua própria documentação técnica:

- `api/README.md`: documentação do backend, endpoints, banco de dados, autenticação e regras da API.
- `webapp/README.md`: documentação da aplicação web, templates, sessão, rotas e integração com a API.

Este README apresenta apenas a visão geral do ecossistema.

## 👨‍💻 Autor

**Jônata Marcelino**

Desenvolvedor de software.

- GitHub: https://github.com/johny83br
- LinkedIn: https://www.linkedin.com/in/jonata-marcelino

---

**DevBook** é um projeto de estudo e portfólio criado para demonstrar, na prática, a construção e integração de uma aplicação web e uma API REST utilizando Go.
