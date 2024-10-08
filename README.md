# Plataforma para gestão de mentorias

Este projeto é uma aplicação open source feita em Go no backend e React no frontend, inicialmente integrada aos serviços do Google Calendar e Gmail, base para discutirmos boas práticas, conceitos e fundamentos.

## Pré-requisitos

- Go versão 1.21 ou superior
- Conta no Google Console para criar credenciais de API
- PostgreSQL instalado (ou acesso a um servidor PostgreSQL)

## Instalação

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/dedevpradev/faladev.git
   cd faladev
   ```

2. **Instale o Go:**

   Certifique-se de ter o Go instalado na versão 1.21 ou superior. Você pode verificar sua versão do Go com o comando:

   ```bash
   go version
   ```

   Para instalar ou atualizar o Go, siga as instruções oficiais: [Instalação do Go](https://golang.org/doc/install)

3. **Configurar Credenciais do Google Console:**

   Para que a aplicação possa acessar o Google Calendar e enviar emails, você precisa configurar as credenciais no Google Console:

   Acesse o Google Cloud Console: [Google Cloud Console](https://console.cloud.google.com/)

   Crie um Novo Projeto:

   - Vá para o painel do Google Cloud Console.
   - Clique em "Select a Project" e depois em "New Project".
   - Dê um nome ao seu projeto e clique em "Create".

   Habilite as APIs Necessárias:

   - Vá para "API & Services" > "Library".
   - Pesquise e habilite a API do Google Calendar.
   - Pesquise e habilite a API do Gmail.

   Configure a Tela de Consentimento OAuth:

   - Vá para "API & Services" > "OAuth consent screen".
   - Escolha "External" e clique em "Create".
   - Preencha as informações necessárias, como nome do aplicativo e email de suporte.
   - Adicione o escopo .../auth/calendar para acesso ao Google Calendar e .../auth/gmail.send para enviar emails.
   - Salve as alterações.

   Crie Credenciais OAuth 2.0:

   - Vá para "API & Services" > "Credentials".
   - Clique em "Create Credentials" e selecione "OAuth 2.0 Client IDs".
   - Escolha "Web application".
   - Adicione os URIs de redirecionamento autorizados. Exemplo: http://localhost:8080/callback
   - Salve as credenciais e anote o Client ID e Client Secret.

4. **Criar um Banco de Dados no PostgreSQL:**
 - Instale o PostgreSQL (ou obtenha acesso a um servidor PostgreSQL): Caso ainda não tenha, siga as instruções oficiais para instalar o PostgreSQL: [Instalação do PostgreSQL](https://www.postgresql.org/download/).

 - Crie um database e anote o nome do banco de dados e do user, senha, host e número da porta.

5. **Configurar Variáveis de Ambiente:**

   Crie um arquivo .env no diretório ./backend e adicione as seguintes variáveis, incluindo suas credenciais do Google:

   ```env
   GOOGLE_MEET_EVENT=https://meet.google.com/link-do-meet-do-evento
   GOOGLE_REDIRECT_URL=http://localhost:8080/callback
   GOOGLE_CLIENT_ID=seu-client-id
   GOOGLE_CLIENT_SECRET=seu-client-secret

   DATABASE_URL=sua-string-de-conexao-postgresql
   ```

   A URL de conexão do PostgreSQL geralmente segue este formato:

   ```
   DATABASE_PORT=
   DATABASE_USER=
   DATABASE_HOST=
   DATABASE_PASSWORD=
   DATABASE_NAME=

   DATABASE_URL="postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}"
   ```

6. **Instalar Dependências:**

   Execute o comando abaixo para instalar as dependências do projeto:

   ```bash
   make deps
   ```

7. **Rodar o Projeto:**

   Para iniciar a aplicação, execute o comando:

   ```bash
   go run cmd/api/main.go
   ```
   Na primeira vez que você rodar o projeto, você deverá acessar o link será gerado no console para autorizar o seu aplicativo na sua conta Google.


8. **Para inspecionar o código:**

   Para verificar a conformidade do código com as diretrizes de estilo, utilize o seguinte comando:

   ```bash
   make lint
   ```

## Instruções para com Docker

   Para iniciar a aplicação, execute o comando:

   ```bash
   docker-compose up -d
   ```

   Para parar e remover contêineres, redes, volumes e imagens usadas pelo docker compose, execute o comando:

   ```bash
   docker-compose down --rmi all
   ```

   Para limpar caches e configurações locais, você pode remover os arquivos de configuração e imagens desnecessárias:

   ```bash
   docker system prune -a --volumes
   ```

## Como usar o Debugger do VScode com Docker?

Tendo iniciado a aplicação com o comando do `docker compose` basta navegar até a aba `Run and Debug` do VSCode.

O comando padrão para acessar essa aba é `Ctrl+Shift+D` ou no Mac `Cmd+Shift+D`. Caso não seja o seu comando padrão você pode se basear [nessa](https://code.visualstudio.com/Docs/editor/debugging) documentação oficial.

Selecione a opção `Connect to server` e clique no botão de Play.
![Run and debug tutorial](.github/images/run_and_debug.png)


## Como atualizar a documentação Swagger?

Instale o pacote golang `swag`:

   ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
   ```

Vá para a pasta backend e rode o seguinte comando:

   ```bash
   path/to/the/go/bin/swag init -g ./cmd/api/main.go -o cmd/docs
   ```

![Run and debug tutorial](.github/images/run_and_debug.png)

## Como Usar

   Para acessar a aplicação:

   ```bash
   http://localhost:3000
   ```

   URL da API:

   ```bash
   http://localhost:8080
   ```

   Para acessar o Jaeger:

   ```bash
   http://localhost:16686/
   ```

   Para acessar o pgAdmin:

   ```bash
   http://localhost:5050/
   ```

   Para acessar a documentação swagger:

   ```bash
   http://localhost:8080/swagger/index.html
   ```

