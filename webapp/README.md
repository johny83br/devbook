# DevBook WebApp

Aplicação web do **DevBook**, desenvolvida em **Go**, responsável pela interface da rede social e pela comunicação com a **DevBook API**.

O WebApp concentra a camada de apresentação, renderiza as páginas HTML no servidor, gerencia a sessão do usuário por meio de cookies seguros e encaminha as operações autenticadas para a API REST.

## 🧩 Sobre o projeto

O **DevBook WebApp** faz parte do ecossistema DevBook e funciona como a camada de interação com o usuário.

Além das funcionalidades de autenticação e publicações, a versão atual também contempla recursos de gerenciamento de usuários e relacionamento entre usuários.

Entre as principais funcionalidades estão:

- Login e logout
- Cadastro de usuários
- Busca de usuários
- Visualização de perfis
- Visualização de seguidores e usuários seguidos
- Seguir e deixar de seguir usuários
- Visualização do perfil do usuário autenticado
- Edição dos dados do usuário
- Atualização de senha
- Exclusão permanente da conta
- Criação de publicações
- Visualização do feed
- Curtir e descurtir publicações
- Atualização de publicações
- Exclusão de publicações
- Navegação com indicação da página ativa
- Alertas e confirmações utilizando SweetAlert2

A aplicação utiliza **Go Templates** para renderização server-side e **JavaScript/jQuery** para interações assíncronas com as rotas do próprio WebApp.

## 🏗️ Arquitetura

O WebApp atua como cliente da DevBook API:

```text
┌─────────────────────┐
│       Usuário       │
└──────────┬──────────┘
           │
           ▼
┌─────────────────────────────┐
│        DevBook WebApp       │
│                             │
│ Go + Templates + JavaScript │
└──────────────┬──────────────┘
               │ HTTP / JSON
               ▼
┌─────────────────────────────┐
│         DevBook API         │
│                             │
│        Go + REST API        │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│            MySQL            │
└─────────────────────────────┘
```

A separação entre WebApp e API permite manter a camada de apresentação independente das regras de negócio e da persistência.

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
- SweetAlert2

### Comunicação

- HTTP
- JSON
- API REST
- AJAX

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
│       ├── main.js
│       ├── usuario.js
│       ├── publicacoes.js
│       ├── jquery.min.js
│       └── bootstrap.min.js
│
├── src/
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   ├── login.go
│   │   ├── logout.go
│   │   ├── paginas.go
│   │   ├── publicacoes.go
│   │   └── usuarios.go
│   ├── cookies/
│   │   └── cookies.go
│   ├── middlewares/
│   │   └── middlewares.go
│   ├── modelos/
│   │   ├── DadosAutenticacao.go
│   │   ├── Publicacao.go
│   │   └── Usuario.go
│   ├── requisicoes/
│   │   └── requisicoes.go
│   ├── respostas/
│   │   └── respostas.go
│   ├── router/
│   │   ├── router.go
│   │   └── rotas/
│   │       ├── home.go
│   │       ├── login.go
│   │       ├── logout.go
│   │       ├── publicacoes.go
│   │       ├── rotas.go
│   │       └── usuarios.go
│   └── utils/
│       └── templates.go
│
└── views/
    ├── login.html
    ├── cadastro.html
    ├── home.html
    ├── perfil.html
    ├── usuario.html
    ├── usuarios.html
    ├── editar-usuario.html
    ├── atualizar-senha.html
    ├── atualizar-publicacao.html
    ├── modal-seguidores.html
    ├── modal-seguindo.html
    │
    └── templates/
        ├── cabecalho.html
        ├── publicacoes.html
        ├── rodape.html
        └── scripts.html
```

## 🔄 Fluxo da aplicação

1. O usuário acessa o WebApp.
2. O servidor Go carrega as configurações do `.env`.
3. Os templates HTML são carregados.
4. O roteador registra as rotas da aplicação.
5. O usuário realiza login ou cadastro.
6. O WebApp encaminha as operações para a DevBook API.
7. Após o login, o JWT retornado pela API é armazenado em um cookie seguro.
8. O middleware verifica a sessão antes de liberar rotas protegidas.
9. Nas operações autenticadas, o token é recuperado do cookie.
10. O token é enviado para a API no header `Authorization`.
11. A API processa a operação e retorna os dados.
12. O WebApp renderiza a resposta ou atualiza a interface via AJAX.

## 🔐 Autenticação e sessão

A autenticação é realizada pela **DevBook API** utilizando JWT.

O WebApp armazena o identificador do usuário e o token em um cookie chamado:

```text
cookie-devbook
```

O cookie é protegido utilizando `gorilla/securecookie` e as chaves:

```env
HASH_KEY="your_secret_hash_key"
BLOCK_KEY="your_secret_block_key"
```

O cookie é configurado com:

```text
HttpOnly: true
Secure: true
Path: /
```

O middleware `Autenticar` verifica o cookie antes de permitir acesso às rotas protegidas.

Quando a sessão não é válida, o usuário é redirecionado para `/login`.

## 🍪 Logout

O logout é realizado através de:

```text
GET /logout
```

O WebApp remove o cookie da sessão e redireciona o usuário para a tela de login.

## 👥 Gerenciamento de usuários

A versão atual do WebApp possui uma camada mais completa de gerenciamento de usuários.

### Busca

O usuário autenticado pode pesquisar outros usuários através de:

```text
GET /buscar-usuarios
```

Os resultados apresentam nome, nickname e data de cadastro.

### Perfil

Cada usuário possui uma página própria:

```text
GET /usuarios/{usuarioId}
```

O perfil apresenta:

- Nome
- Nick
- Data de cadastro
- Quantidade de seguidores
- Quantidade de usuários seguidos
- Publicações
- Ação para seguir ou deixar de seguir

### Perfil do usuário autenticado

```text
GET /perfil
```

O perfil próprio permite consultar:

- Dados pessoais
- Seguidores
- Usuários seguidos
- Publicações
- Edição de dados
- Atualização de senha
- Exclusão da conta

### Seguidores

O perfil apresenta modais para consultar:

- Seguidores
- Usuários que o perfil está seguindo

O relacionamento pode ser alterado através das ações:

```text
POST /usuarios/{usuarioId}/seguir
POST /usuarios/{usuarioId}/parar-de-seguir
```

## ✏️ Edição de usuário

O WebApp disponibiliza uma página específica para edição dos dados:

```text
GET /editar-usuario
```

A atualização é enviada para:

```text
PUT /editar-usuario
```

Os dados tratados são:

- Nome
- E-mail
- Nick

## 🔑 Atualização de senha

A aplicação possui uma tela específica para alteração da senha:

```text
GET /atualizar-senha
```

A atualização é realizada através de:

```text
POST /atualizar-senha
```

Antes de enviar a requisição, o JavaScript verifica se a nova senha e a confirmação são iguais.

A validação final da senha continua sendo responsabilidade da API.

## 🗑️ Exclusão de conta

O usuário autenticado pode solicitar a exclusão permanente da própria conta:

```text
DELETE /deletar-usuario
```

A interface apresenta uma confirmação antes de executar a operação.

Após a exclusão, o WebApp realiza o logout e retorna à tela de login.

## 📝 Publicações

O WebApp permite criar e gerenciar publicações.

### Criar

```text
POST /publicacoes
```

### Curtir

```text
POST /publicacoes/{publicacaoId}/curtir
```

### Descurtir

```text
POST /publicacoes/{publicacaoId}/descurtir
```

### Atualizar

A tela de edição é carregada por:

```text
GET /publicacoes/{publicacaoId}/atualizar
```

A alteração é enviada para:

```text
PUT /publicacoes/{publicacaoId}
```

### Excluir

```text
DELETE /publicacoes/{publicacaoId}
```

As operações de publicação utilizam AJAX, permitindo atualizar o estado da interface sem depender de um recarregamento completo da página.

## 🌐 Rotas

### Autenticação

| Método | Rota | Auth | Descrição |
|---|---|---:|---|
| `GET` | `/` | Não | Tela de login |
| `GET` | `/login` | Não | Tela de login |
| `POST` | `/login` | Não | Autenticação |
| `GET` | `/logout` | Sim | Encerra a sessão |

### Usuários

| Método | Rota | Auth | Descrição |
|---|---|---:|---|
| `GET` | `/criar-usuario` | Não | Tela de cadastro |
| `POST` | `/criar-usuario` | Não | Cria usuário |
| `GET` | `/buscar-usuarios` | Sim | Pesquisa usuários |
| `GET` | `/usuarios/{usuarioId}` | Sim | Visualiza perfil |
| `POST` | `/usuarios/{usuarioId}/seguir` | Sim | Segue usuário |
| `POST` | `/usuarios/{usuarioId}/parar-de-seguir` | Sim | Deixa de seguir |
| `GET` | `/perfil` | Sim | Perfil do usuário autenticado |
| `GET` | `/editar-usuario` | Sim | Tela de edição |
| `PUT` | `/editar-usuario` | Sim | Atualiza dados |
| `GET` | `/atualizar-senha` | Sim | Tela de alteração |
| `POST` | `/atualizar-senha` | Sim | Atualiza senha |
| `DELETE` | `/deletar-usuario` | Sim | Exclui conta |

### Publicações

| Método | Rota | Auth | Descrição |
|---|---|---:|---|
| `POST` | `/publicacoes` | Sim | Cria publicação |
| `POST` | `/publicacoes/{publicacaoId}/curtir` | Sim | Curte publicação |
| `POST` | `/publicacoes/{publicacaoId}/descurtir` | Sim | Remove curtida |
| `GET` | `/publicacoes/{publicacaoId}/atualizar` | Sim | Tela de edição |
| `PUT` | `/publicacoes/{publicacaoId}` | Sim | Atualiza publicação |
| `DELETE` | `/publicacoes/{publicacaoId}` | Sim | Exclui publicação |

### Arquivos estáticos

Os arquivos de `assets/` são disponibilizados através de:

```text
/assets/
```

## ⚡ JavaScript e AJAX

O frontend utiliza JavaScript e jQuery para realizar operações assíncronas.

### `login.js`

Responsável por:

- Validar e-mail
- Validar senha
- Enviar login
- Redirecionar para `/home`

### `cadastro.js`

Responsável por:

- Validar confirmação da senha
- Enviar cadastro
- Exibir mensagens de sucesso e erro
- Redirecionar para o login

### `usuario.js`

Centraliza as interações relacionadas aos usuários:

- Seguir
- Deixar de seguir
- Editar dados
- Atualizar senha
- Excluir conta

### `publicacoes.js`

Gerencia:

- Criação de publicações
- Curtidas
- Descurtidas
- Atualização de publicações
- Exclusão de publicações

### `main.js`

Responsável por identificar a rota atual e destacar o item correspondente no menu de navegação.

## 🔔 Feedback da interface

A versão atual utiliza **SweetAlert2** para apresentar mensagens de:

- Sucesso
- Erro
- Avisos
- Confirmação de ações irreversíveis

Isso é utilizado especialmente em operações como:

- Login
- Cadastro
- Seguir usuários
- Atualizar dados
- Atualizar senha
- Excluir conta
- Criar publicações
- Atualizar publicações
- Excluir publicações

## 🖥️ Templates

A aplicação utiliza `html/template` para renderização server-side.

Além das páginas principais, existem templates reutilizáveis para:

- Cabeçalho
- Rodapé
- Scripts
- Publicações
- Modal de seguidores
- Modal de usuários seguidos

Essa estrutura reduz duplicação de HTML e mantém componentes comuns centralizados.

## 🚀 Comunicação com a API

O WebApp utiliza:

```env
API_URL="http://localhost"
API_PORT=5000
```

As requisições autenticadas recuperam o JWT do cookie e adicionam:

```http
Authorization: Bearer SEU_TOKEN
```

A API continua sendo responsável pelas regras de negócio, validações e persistência.

O WebApp atua como uma camada intermediária entre o navegador e a API.

## ⚙️ Configuração

Crie um `.env` baseado no `example.env`:

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

| Variável | Descrição | Exemplo |
|---|---|---|
| `API_URL` | URL base da API | `http://localhost` |
| `API_PORT` | Porta da API | `5000` |
| `APP_PORT` | Porta do WebApp | `4000` |
| `HASH_KEY` | Chave de assinatura dos cookies | chave secreta |
| `BLOCK_KEY` | Chave de proteção dos cookies | chave secreta |

> Não versione credenciais ou chaves reais. O arquivo `.env` deve permanecer fora do controle de versão.

## 🚀 Executando o projeto

### Pré-requisitos

- Go 1.25.2 ou compatível
- DevBook API em execução
- Git

### Instalação

```bash
git clone https://github.com/johny83br/devbook.git
cd devbook/webapp
go mod download
```

Configure o ambiente:

```bash
cp example.env .env
```

Depois execute:

```bash
go run .
```

O WebApp será executado na porta definida em:

```env
APP_PORT=4000
```

Acesse:

```text
http://localhost:4000
```

## 🔗 Dependência da API

Para utilizar o ecossistema completo, a **DevBook API** deve estar disponível.

Uma configuração local típica é:

```text
WebApp
http://localhost:4000
       │
       │ HTTP / JSON
       ▼
API
http://localhost:5000
       │
       ▼
MySQL
```

## 🎯 Conceitos demonstrados

O projeto demonstra conceitos importantes de desenvolvimento web com Go:

- Desenvolvimento de aplicações web com `net/http`
- Roteamento com Gorilla Mux
- Arquitetura cliente-servidor
- Separação entre WebApp e API
- Renderização server-side com Go Templates
- Controllers e middlewares
- Consumo de APIs REST
- Comunicação HTTP/JSON
- Autenticação baseada em JWT
- Gerenciamento de sessão com cookies
- Proteção de cookies com SecureCookie
- Operações assíncronas com AJAX
- Integração entre Go e JavaScript
- CRUD de usuários e publicações
- Relacionamento entre usuários
- Uso de variáveis de ambiente
- Organização de código em pacotes
- Uso de concorrência com goroutines para carregar dados de perfil em paralelo

## 📚 Relação com o DevBook API

O WebApp e a API possuem responsabilidades distintas:

| Componente | Responsabilidade |
|---|---|
| `webapp` | Interface, navegação, sessão, templates e interação com o usuário |
| `api` | Regras de negócio, autenticação, usuários, publicações e persistência |
| `MySQL` | Armazenamento dos dados |

Essa separação permite que o frontend web consuma os recursos da API sem acessar diretamente o banco de dados.

## 👨‍💻 Autor

**Jônata Marcelino**

Desenvolvedor de software.

- GitHub: https://github.com/johny83br
- LinkedIn: https://www.linkedin.com/in/jonata-marcelino

---

Projeto desenvolvido para estudos e prática de desenvolvimento web com Go, integração com APIs REST, autenticação, gerenciamento de sessão e arquitetura cliente-servidor.
