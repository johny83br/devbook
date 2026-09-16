# DevBook WebApp

Aplicação web do **DevBook**, desenvolvida em **Go**, responsável pela interface e pela camada de comunicação entre o usuário e a **DevBook API**.

Enquanto a API concentra as regras de negócio e o acesso ao banco de dados, este projeto funciona como a aplicação web que apresenta as páginas da plataforma, processa formulários, mantém a sessão do usuário e encaminha as operações para a API REST.

## 🧩 Sobre o projeto

O **DevBook WebApp** é a camada web de uma aplicação no estilo rede social para desenvolvedores.

A aplicação permite que o usuário:

- Acesse a tela de login
- Crie uma nova conta
- Faça autenticação
- Acesse a página principal
- Visualize publicações
- Crie novas publicações
- Curta publicações
- Remova curtidas

A aplicação utiliza templates HTML renderizados no servidor e JavaScript/jQuery para realizar operações assíncronas, principalmente na criação e interação com publicações.

## 🏗️ Arquitetura

O projeto funciona como um cliente web para a **DevBook API**:

```text
┌─────────────────────┐
│       Usuário       │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│    DevBook WebApp   │
│                     │
│ Go + HTML Templates │
│ JavaScript + jQuery │
└──────────┬──────────┘
           │ HTTP / JSON
           ▼
┌─────────────────────┐
│     DevBook API     │
│                     │
│       Go REST       │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────┐
│       MySQL         │
└─────────────────────┘
```

Essa separação permite que a interface web e a API sejam executadas como aplicações independentes.

## 🛠️ Tecnologias utilizadas

### Backend Web

- **Go 1.25.2**
- **Gorilla Mux**
- **GoDotEnv**
- **Gorilla SecureCookie**
- `html/template`
- `net/http`

### Frontend

- HTML5
- CSS3
- JavaScript
- jQuery
- Bootstrap
- Font Awesome

### Comunicação

- HTTP
- JSON
- REST API

## 📁 Estrutura do projeto

```text
webapp/
├── main.go
├── go.mod
├── go.sum
├── example.env
├── .env
│
├── assets/
│   ├── css/
│   │   ├── bootstrap.min.css
│   │   └── login.css
│   │
│   └── js/
│       ├── cadastro.js
│       ├── login.js
│       ├── publicacoes.js
│       ├── jquery.min.js
│       └── bootstrap.min.js
│
├── src/
│   ├── config/
│   │   └── config.go
│   │
│   ├── controllers/
│   │   ├── login.go
│   │   ├── paginas.go
│   │   ├── publicacoes.go
│   │   └── usuarios.go
│   │
│   ├── cookies/
│   │   └── cookies.go
│   │
│   ├── middlewares/
│   │   └── middlewares.go
│   │
│   ├── modelos/
│   │   ├── DadosAutenticacao.go
│   │   └── Publicacao.go
│   │
│   ├── requisicoes/
│   │   └── requisicoes.go
│   │
│   ├── respostas/
│   │   └── respostas.go
│   │
│   ├── router/
│   │   ├── router.go
│   │   └── rotas/
│   │       ├── home.go
│   │       ├── login.go
│   │       ├── publicacoes.go
│   │       ├── rotas.go
│   │       └── usuarios.go
│   │
│   └── utils/
│       └── templates.go
│
└── views/
    ├── login.html
    ├── cadastro.html
    ├── home.html
    │
    └── templates/
        ├── cabecalho.html
        ├── publicacoes.html
        ├── rodape.html
        └── scripts.html
```

## 🔄 Fluxo da aplicação

O fluxo principal da aplicação funciona da seguinte maneira:

1. O usuário acessa o WebApp.
2. O servidor Go carrega as configurações do ambiente.
3. Os templates HTML são carregados.
4. O roteador registra as rotas da aplicação.
5. O usuário realiza o login ou cadastro.
6. O WebApp se comunica com a DevBook API através de requisições HTTP.
7. Após o login, o token recebido pela API é armazenado em um cookie seguro.
8. Nas operações autenticadas, o WebApp recupera o token do cookie.
9. O token é enviado para a API no header `Authorization`.
10. A resposta da API é processada pelo WebApp e apresentada ao usuário.

## 🔐 Autenticação e sessão

O WebApp não realiza diretamente a autenticação contra o banco de dados.

Quando o usuário faz login:

```text
Browser
   │
   │ POST /login
   ▼
WebApp
   │
   │ POST /login
   ▼
DevBook API
   │
   │ JWT
   ▼
WebApp
   │
   │ Cookie seguro
   ▼
Browser
```

O token JWT retornado pela API é armazenado dentro do cookie `cookie-devbook`.

Para proteger o conteúdo da sessão, o projeto utiliza:

- `gorilla/securecookie`
- `HASH_KEY`
- `BLOCK_KEY`
- Cookie `HttpOnly`
- Cookie `Secure`

O middleware `Autenticar` verifica a existência e validade do cookie antes de permitir o acesso às páginas protegidas.

Quando a autenticação não está disponível, o usuário é redirecionado para:

```text
/login
```

## 🍪 Cookies

A sessão do usuário é armazenada em um cookie chamado:

```text
cookie-devbook
```

O cookie contém:

```json
{
  "id": "ID_DO_USUARIO",
  "token": "JWT"
}
```

Os dados são codificados utilizando `securecookie`.

As principais propriedades configuradas no cookie são:

```text
HttpOnly: true
Secure: true
Path: /
```

## 🌐 Rotas

### Login

| Método | Rota | Autenticação | Descrição |
|---|---|---|---|
| `GET` | `/` | Não | Exibe a tela de login |
| `GET` | `/login` | Não | Exibe a tela de login |
| `POST` | `/login` | Não | Realiza o login através da API |

### Cadastro

| Método | Rota | Autenticação | Descrição |
|---|---|---|---|
| `GET` | `/criar-usuario` | Não | Exibe a tela de cadastro |
| `POST` | `/criar-usuario` | Não | Cria um usuário através da API |

### Página principal

| Método | Rota | Autenticação | Descrição |
|---|---|---|---|
| `GET` | `/home` | Sim | Exibe o feed de publicações |

### Publicações

| Método | Rota | Autenticação | Descrição |
|---|---|---|---|
| `POST` | `/publicacoes` | Sim | Cria uma publicação |
| `POST` | `/publicacoes/{publicacaoId}/curtir` | Sim | Curte uma publicação |
| `POST` | `/publicacoes/{publicacaoId}/descurtir` | Sim | Remove uma curtida |

### Arquivos estáticos

Os arquivos presentes em `assets/` são disponibilizados através da rota:

```text
/assets/
```

Exemplos:

```text
/assets/css/login.css
/assets/js/login.js
/assets/js/publicacoes.js
```

## 📡 Comunicação com a API

O WebApp utiliza a variável `API_URL` e a porta definida em `API_PORT` para localizar a DevBook API.

Por padrão:

```text
API_URL=http://localhost
API_PORT=5000
```

A aplicação web, por sua vez, utiliza:

```text
APP_PORT=4000
```

Assim, em um ambiente local, a arquitetura pode ser executada como:

```text
WebApp
http://localhost:4000

        │
        ▼

API
http://localhost:5000
```

## 🔁 Requisições autenticadas

As requisições que exigem autenticação são realizadas pela função:

```text
FazerRequisicaoComAutenticacao
```

Ela recupera o token armazenado no cookie e adiciona o header:

```http
Authorization: Bearer SEU_TOKEN
```

Isso permite que a API identifique o usuário responsável pela operação.

## 🖥️ Templates

O projeto utiliza o pacote padrão `html/template` do Go.

Os templates principais são:

- `login.html`
- `cadastro.html`
- `home.html`

Também existem templates reutilizáveis:

- Cabeçalho
- Rodapé
- Scripts
- Estrutura das publicações

A aplicação carrega os templates utilizando:

```go
template.ParseGlob("views/*.html")
template.ParseGlob("views/templates/*.html")
```

Isso permite reutilizar componentes HTML entre diferentes páginas.

## ⚡ Interações com JavaScript

A interface utiliza jQuery para realizar chamadas assíncronas.

### Login

O arquivo:

```text
assets/js/login.js
```

valida os campos de e-mail e senha e envia os dados para:

```text
POST /login
```

Após o login, o usuário é direcionado para:

```text
/home
```

### Cadastro

O arquivo:

```text
assets/js/cadastro.js
```

realiza a validação da confirmação da senha e envia os dados para:

```text
POST /criar-usuario
```

Após o cadastro, o usuário é direcionado para:

```text
/login
```

### Publicações

O arquivo:

```text
assets/js/publicacoes.js
```

é responsável pelas interações relacionadas às publicações.

Ele permite:

- Criar publicações
- Curtir publicações
- Remover curtidas
- Atualizar visualmente o contador de curtidas

As operações são realizadas sem a necessidade de recarregar manualmente a página para cada ação.

## 🎨 Interface

A interface utiliza:

- **Bootstrap** para estrutura e componentes
- **Font Awesome** para ícones
- CSS próprio para as telas de autenticação
- Templates do Go para renderização no servidor

A página inicial apresenta:

- Barra de navegação
- Formulário de nova publicação
- Lista de publicações
- Informações do autor
- Data da publicação
- Contador de curtidas
- Ações de curtida

## ⚙️ Configuração

O projeto utiliza variáveis de ambiente.

Crie um arquivo `.env` baseado no `example.env`:

```bash
cp example.env .env
```

Configure:

```env
API_URL="http://localhost"
API_PORT=5000
APP_PORT=4000
HASH_KEY="your_secret_hash_key"
BLOCK_KEY="your_secret_block_key"
```

### Variáveis

| Variável | Descrição | Exemplo |
|---|---|---|
| `API_URL` | URL base da DevBook API | `http://localhost` |
| `API_PORT` | Porta da API | `5000` |
| `APP_PORT` | Porta do WebApp | `4000` |
| `HASH_KEY` | Chave para assinatura dos cookies | chave secreta |
| `BLOCK_KEY` | Chave para criptografia dos cookies | chave secreta |

> Nunca versione chaves reais ou credenciais no repositório. O arquivo `.env` deve permanecer fora do controle de versão.

## 🚀 Executando o projeto

### Pré-requisitos

- Go 1.25.2 ou compatível
- DevBook API em execução
- Git

### Clone o projeto

```bash
git clone https://github.com/johny83br/devbook.git
```

Acesse o diretório do WebApp:

```bash
cd devbook/webapp
```

Instale as dependências:

```bash
go mod download
```

Configure o arquivo `.env`:

```bash
cp example.env .env
```

Ajuste as configurações conforme o ambiente.

### Execute

```bash
go run .
```

O WebApp será executado na porta configurada em:

```env
APP_PORT=4000
```

Acesse:

```text
http://localhost:4000
```

## 🔗 Dependência da API

O WebApp depende da **DevBook API** para realizar operações de autenticação, usuários e publicações.

Portanto, para utilizar a aplicação completa, a API deve estar disponível no endereço configurado em:

```env
API_URL
API_PORT
```

Exemplo:

```text
WebApp → http://localhost:4000
API    → http://localhost:5000
```

## 🎯 Objetivos do projeto

O projeto demonstra, na prática, conceitos importantes de desenvolvimento web com Go:

- Criação de aplicações web com `net/http`
- Roteamento com Gorilla Mux
- Renderização de HTML no servidor
- Organização de controllers
- Middlewares
- Autenticação baseada em JWT
- Gerenciamento de sessão através de cookies
- Cookies protegidos com `securecookie`
- Comunicação entre aplicações através de HTTP
- Consumo de APIs REST
- Uso de JSON
- Requisições AJAX
- Integração entre Go e JavaScript
- Separação entre frontend web e backend API
- Configuração através de variáveis de ambiente

## 📚 Relação com o DevBook API

Este projeto faz parte do ecossistema **DevBook** e funciona em conjunto com a API.

A divisão de responsabilidades é:

| Projeto | Responsabilidade |
|---|---|
| `webapp` | Interface web, sessão, templates e interação com o usuário |
| `api` | Regras de negócio, autenticação, usuários, publicações e persistência |
| `MySQL` | Armazenamento dos dados |

Essa estrutura permite separar a camada de apresentação da camada responsável pelos dados e regras da aplicação.

## 👨‍💻 Autor

**Jônata Marcelino**

Desenvolvedor de software.

- GitHub: https://github.com/johny83br
- LinkedIn: https://www.linkedin.com/in/jonata-marcelino

---

Projeto desenvolvido para estudos e prática de desenvolvimento web com Go, integração com APIs REST, autenticação e arquitetura em camadas.
