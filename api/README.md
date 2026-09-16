# DevBook API

API REST desenvolvida em **Go** para o backend da aplicação **DevBook**, uma rede social voltada para publicação e interação entre usuários.

O projeto disponibiliza recursos para cadastro e autenticação de usuários, gerenciamento de perfis, relacionamento entre seguidores e publicações com curtidas.

## Sobre o projeto

O **DevBook API** é o backend de uma aplicação de rede social. A API é responsável por receber as requisições do cliente, validar os dados, aplicar as regras de negócio, acessar o banco de dados e retornar as respostas em JSON.

Entre as principais funcionalidades estão:

- Cadastro de usuários
- Login e autenticação
- Geração e validação de tokens JWT
- Consulta e atualização de usuários
- Exclusão de usuários
- Alteração de senha
- Sistema de seguidores
- Listagem de seguidores e usuários seguidos
- Criação, consulta, edição e exclusão de publicações
- Feed de publicações baseado nos usuários seguidos
- Curtidas e remoção de curtidas
- Validação e tratamento de erros

## Tecnologias utilizadas

- **Go 1.25.2**
- **MySQL**
- **Gorilla Mux** para roteamento HTTP
- **JWT** para autenticação
- **bcrypt** para hash e validação de senhas
- **GoDotEnv** para carregamento das variáveis de ambiente
- **Checkmail** para validação do formato de e-mail
- **database/sql** para acesso ao banco de dados

## Arquitetura

O projeto está organizado em diferentes pacotes, separando responsabilidades da aplicação:

```text
api/
├── main.go
├── go.mod
├── go.sum
├── .env.example
├── sql/
│   ├── sql.sql
│   └── dados.sql
└── src/
    ├── autenticacao/
    │   └── token.go
    ├── banco/
    │   └── banco.go
    ├── config/
    │   └── config.go
    ├── controllers/
    │   ├── login.go
    │   ├── publicacoes.go
    │   └── usuarios.go
    ├── middlewares/
    │   └── middlewares.go
    ├── modelos/
    │   ├── DadosAutenticacao.go
    │   ├── Publicacao.go
    │   ├── Senha.go
    │   └── Usuario.go
    ├── repositorios/
    │   ├── publicacoes.go
    │   └── usuarios.go
    ├── respostas/
    │   └── respostas.go
    ├── router/
    │   ├── router.go
    │   └── rotas/
    │       ├── login.go
    │       ├── publicacoes.go
    │       ├── rotas.go
    │       └── usuarios.go
    └── seguranca/
        └── seguranca.go
```

A aplicação segue uma separação simples entre:

- **Controllers**: recebem e processam as requisições HTTP.
- **Modelos**: representam os dados e concentram validações.
- **Repositórios**: encapsulam o acesso ao banco de dados.
- **Router/Rotas**: define os endpoints e métodos HTTP.
- **Middlewares**: executam funcionalidades comuns, como logging e autenticação.
- **Autenticação**: criação, validação e extração de informações dos tokens JWT.
- **Segurança**: geração e validação de hashes de senha.
- **Banco**: estabelece a conexão com o MySQL.
- **Respostas**: padroniza o retorno de dados e erros em JSON.
- **Configuração**: carrega as configurações da aplicação a partir de variáveis de ambiente.

## Autenticação

A autenticação utiliza **JWT (JSON Web Token)** com o algoritmo HMAC SHA-256.

Após o login, a API retorna um token com:

- ID do usuário
- Indicação de autorização
- Data de expiração
- Identificador do usuário

O token possui validade de **6 horas**.

Endpoints protegidos devem receber o token no header:

```http
Authorization: Bearer SEU_TOKEN
```

O middleware de autenticação valida o token antes de encaminhar a requisição ao controller.

## Segurança de senhas

As senhas não são armazenadas diretamente no banco.

Durante o cadastro ou alteração da senha, o projeto utiliza **bcrypt** para gerar o hash. No login, a senha informada é comparada com o hash armazenado.

## Banco de dados

O projeto utiliza **MySQL** e possui três tabelas principais:

### `usuarios`

Armazena os dados dos usuários:

- ID
- Nome
- Nick
- E-mail
- Senha
- Data de criação

### `seguidores`

Representa o relacionamento entre usuários, permitindo que um usuário siga outro.

### `publicacoes`

Armazena as publicações realizadas pelos usuários:

- ID
- Título
- Conteúdo
- Autor
- Quantidade de curtidas
- Data de criação

O relacionamento entre as tabelas utiliza chaves estrangeiras e exclusão em cascata para usuários e suas relações/publicações.

Os scripts de criação e dados iniciais estão disponíveis no diretório:

```text
sql/
```

## Endpoints

### Autenticação

| Método | Endpoint | Autenticação | Descrição |
|---|---|---|---|
| `POST` | `/login` | Não | Realiza o login do usuário |

### Usuários

| Método | Endpoint | Autenticação | Descrição |
|---|---|---|---|
| `POST` | `/api/usuarios` | Não | Cadastra um usuário |
| `GET` | `/api/usuarios` | Sim | Busca usuários por nome ou nickname |
| `GET` | `/api/usuarios/{usuarioId}` | Sim | Busca um usuário pelo ID |
| `PUT` | `/api/usuarios/{usuarioId}` | Sim | Atualiza os dados do próprio usuário |
| `DELETE` | `/api/usuarios/{usuarioId}` | Sim | Remove o próprio usuário |
| `POST` | `/api/usuarios/{usuarioId}/seguir` | Sim | Segue um usuário |
| `POST` | `/api/usuarios/{usuarioId}/parar-de-seguir` | Sim | Deixa de seguir um usuário |
| `GET` | `/api/usuarios/{usuarioId}/seguidores` | Sim | Lista os seguidores |
| `GET` | `/api/usuarios/{usuarioId}/seguindo` | Sim | Lista os usuários seguidos |
| `POST` | `/api/usuarios/{usuarioId}/atualizar-senha` | Sim | Atualiza a senha |

### Publicações

| Método | Endpoint | Autenticação | Descrição |
|---|---|---|---|
| `POST` | `/api/publicacoes` | Sim | Cria uma publicação |
| `GET` | `/api/publicacoes` | Sim | Lista as publicações do feed |
| `GET` | `/api/publicacoes/{publicacaoId}` | Sim | Busca uma publicação pelo ID |
| `PUT` | `/api/publicacoes/{publicacaoId}` | Sim | Atualiza uma publicação do próprio usuário |
| `DELETE` | `/api/publicacoes/{publicacaoId}` | Sim | Remove uma publicação do próprio usuário |
| `POST` | `/api/publicacoes/{publicacaoId}/curtir` | Sim | Curte uma publicação |
| `POST` | `/api/publicacoes/{publicacaoId}/descurtir` | Sim | Remove a curtida |
| `GET` | `/api/usuarios/{usuarioId}/publicacoes` | Sim | Lista as publicações de um usuário |

## Exemplos de requisições

### Cadastro de usuário

```http
POST /api/usuarios
Content-Type: application/json
```

```json
{
  "nome": "Jônata Marcelino",
  "nick": "jonata",
  "email": "jonata@example.com",
  "senha": "123456"
}
```

### Login

```http
POST /login
Content-Type: application/json
```

```json
{
  "email": "jonata@example.com",
  "senha": "123456"
}
```

Resposta:

```json
{
  "id": "1",
  "token": "SEU_TOKEN_JWT"
}
```

### Criar publicação

```http
POST /api/publicacoes
Authorization: Bearer SEU_TOKEN
Content-Type: application/json
```

```json
{
  "titulo": "Minha primeira publicação",
  "conteudo": "Olá, DevBook!"
}
```

### Curtir publicação

```http
POST /api/publicacoes/1/curtir
Authorization: Bearer SEU_TOKEN
```

## Configuração do ambiente

O projeto utiliza variáveis de ambiente para configurar a porta da API, conexão com o banco de dados e chave utilizada na assinatura dos tokens.

Crie um arquivo `.env` baseado no `.env.example`:

```bash
cp .env.example .env
```

Configure as variáveis de acordo com o seu ambiente:

```env
API_PORTA=9000

DB_USUARIO=seu_usuario
DB_SENHA=sua_senha
DB_HOST=localhost
DB_PORTA=3306
DB_NOME=devbook

SECRET_KEY=sua_chave_secreta
```

> Não versionar o arquivo `.env`, especialmente quando ele contém credenciais ou chaves privadas.

## Pré-requisitos

Para executar o projeto localmente, você precisará ter instalado:

- Go 1.25.2 ou compatível
- MySQL
- Git

## Instalação

Clone o repositório:

```bash
git clone https://github.com/johny83br/devbook.git
```

Acesse o diretório da API:

```bash
cd devbook/api
```

Instale as dependências:

```bash
go mod download
```

Configure o banco de dados executando os scripts SQL disponíveis em:

```text
sql/sql.sql
sql/dados.sql
```

Configure o arquivo `.env` e execute a aplicação:

```bash
go run .
```

A API será iniciada na porta configurada em `API_PORTA`. Caso essa variável não seja convertida corretamente, a aplicação utiliza a porta `9000` como fallback.

## Executando o binário

Também é possível gerar um executável:

```bash
go build -o devbook-api
```

Depois:

```bash
./devbook-api
```

No Windows:

```bash
devbook-api.exe
```

## Tratamento de erros

As respostas de erro são padronizadas em JSON:

```json
{
  "error": "Mensagem do erro"
}
```

A API utiliza códigos HTTP para indicar o resultado da operação, incluindo:

- `200 OK`
- `201 Created`
- `204 No Content`
- `400 Bad Request`
- `401 Unauthorized`
- `403 Forbidden`
- `422 Unprocessable Entity`
- `500 Internal Server Error`

## Middleware

O projeto possui dois middlewares principais:

### Logger

Registra informações básicas das requisições recebidas:

```text
Método URI Host
```

### Autenticação

Valida o token JWT enviado no header `Authorization` e impede o acesso aos endpoints protegidos quando o token é inválido ou ausente.

## Objetivo do projeto

O DevBook API é um projeto prático para desenvolvimento de uma API REST utilizando Go, explorando conceitos como:

- Desenvolvimento de APIs HTTP
- Organização de aplicações Go em pacotes
- Roteamento com Gorilla Mux
- Persistência de dados com MySQL
- Autenticação baseada em JWT
- Hash de senhas com bcrypt
- Middlewares
- Validação de dados
- Separação entre controllers e repositórios
- Relacionamentos entre usuários
- Operações CRUD
- Tratamento padronizado de respostas HTTP

## Autor

**Jônata Marcelino**

Desenvolvedor de software.

- GitHub: https://github.com/johny83br
- LinkedIn: https://www.linkedin.com/in/jonata-marcelino

---

Projeto desenvolvido para estudos e prática de desenvolvimento backend com Go e APIs REST.
