FROM golang:1.18-alpine

WORKDIR /app

# Copie go.mod e go.sum para o diretório de trabalho
COPY cmd/go.mod cmd/go.sum ./cmd/

# Copie todo o código da aplicação para /app
COPY . .

# Defina o diretório de trabalho para /app/cmd
WORKDIR /app/cmd

# Baixe as dependências
RUN go mod download

# Construa a aplicação
RUN go build -o main ./main.go

EXPOSE 8080
CMD ["./main"]