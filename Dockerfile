# Usando a imagem oficial do Golang
FROM golang:1.22-alpine

# Definir o diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar as dependências do Go
RUN go mod download

# Copiar o restante do código para o contêiner
COPY . .

# Compilar o binário do Go
RUN go build -o main ./cmd/main.go  # Verifique o caminho para o main.go

# Expor a porta 8080
EXPOSE 8080

# Comando para rodar o app Go
CMD ["./main"]