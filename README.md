# api-dev-book

ste projeto é baseado no curso da Udemy: "Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!" O projeto apresenta algumas diferenças em relação ao original, sendo a principal delas o banco de dados, que foi adaptado para utilizar o PostgreSQL, ao contrário do projeto original que utiliza o MySQL. Além disso, foram adicionados arquivos para possibilitar a criação e a execução via contêiner Docker. Observação: Apenas a parte do backend foi desenvolvida.

## Links do Projeto original

- [Curso na Udemy] - https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/
- [Repositório Original] - https://github.com/OtavioGallego/devbook 

## Descrição
Usuários podem se cadastrar na aplicação, e os usuários cadastrados podem criar publicações e seguir outros usuários. Além disso, os usuários podem curtir ou descurtir publicações, sem limites de curtidas por usuário.

## Documentação do Projeto

### Tecnologias usadas no projeto

- Postgres
- jwt-go
- gorilla/mux
- badoux/checkmail

### Instruções para Execução

Clone o repositório: 
```bash
$ git clone https://github.com/Golang-pratics/api-dev-book.git
```

### Execução Local:

```bash
$ go run main.go
```

### Execução via docker-compose:

```bash
$ docker-compose up -d
Obs: Isso realizará o build da imagem e executará o projeto.
```
Os scripts sql para criação das tabelas estão no diretório sql/sql.sql

### Variáveis de Ambiente

DB_HOST: 172.18.0.1 (endereço do banco de dados, padrão do contêiner Docker)

DB_USUARIO: postgres (usuário do banco)

DB_SENHA: postgres (senha do banco)

DB_NOME: devbook (nome do banco)

DB_PORT: 5432 (porta do banco)

API_PORTA: 5000 (porta de execução da API)

SECRET_KEY: secret (variável utilizada para a criptografia do token)



### Futuras implementações:

- Documentação Swagger
- Envio de notificações por email, seja para curtidas ou quando um usuário começar a seguir outro usuário.
- Cache com Redis
- Mensageria com Kafka: serviço responsável pelo envio de emails.