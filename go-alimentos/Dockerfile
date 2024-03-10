# Stage 1: Build Stage
FROM golang:1.22.1 as build
WORKDIR /app

# Copie os arquivos de configuração e baixe as dependências antes de copiar o restante do código-fonte
COPY go.* ./
RUN go mod download

# Copie o restante do código-fonte
COPY . .

# Compile o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o go-alimentos .

# Stage 2: Final Stage
FROM alpine:latest
WORKDIR /app

# Copie apenas os arquivos necessários do estágio de compilação
COPY --from=build /app/go-alimentos .
COPY tabela_alimentos.json .

# Configure as variáveis de ambiente, se necessário
ENV PORT=8080

# Exponha a porta
EXPOSE $PORT

# Comando padrão para iniciar o aplicativo
CMD ["./go-alimentos"]
