# Use a imagem oficial do Golang como base
FROM golang:1.17

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# variaveis de ambiente
# ENV DB_HOST=database

# Copia o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa e instala as dependências
RUN go mod download

# Copia o resto do código-fonte para o diretório de trabalho
COPY . .

# Compila o aplicativo
RUN go build -o app

# Comando padrão para executar o aplicativo quando o contêiner iniciar
CMD ["./app"]