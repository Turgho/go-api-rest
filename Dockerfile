# Usar a imagem oficial do Golang como base
FROM golang:1.22

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos do projeto para o diretório de trabalho
COPY go.mod .
COPY go.sum .

# Baixar as dependências
RUN go mod download

# Copiar o restante dos arquivos do projeto
COPY . .

# Expor a porta que a aplicação usará
EXPOSE 5050

# Compilar o binário
RUN go build -o main cmd/main.go

# Comando para rodar a aplicação
CMD [ "./main" ]