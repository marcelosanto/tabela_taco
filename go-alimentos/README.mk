# Tabela de Alimentos API em Golang

Esta é uma API simples em Golang para acessar uma tabela de alimentos. A tabela é carregada a partir de um arquivo JSON local e fornece endpoints para obter todos os alimentos ou buscar alimentos por nome.

## Pré-requisitos

Certifique-se de ter o Go instalado em seu ambiente de desenvolvimento.

## Instalação e Execução

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```


2.  Execute o servidor:

```go
go run main.go
```

O servidor será iniciado em http://localhost:8080.

## Endpoints

### 1. Obter todos os alimentos

```bash
`GET /tabela-alimentos` 
```

Retorna todos os alimentos na tabela.

### 2. Buscar alimentos por nome


`GET /buscar-alimento?nome=<nome-do-alimento>` 

Retorna alimentos que contenham o nome especificado (case-insensitive).

# Usando Docker

Certifique-se de ter o Docker instalado em sua máquina.

### Construa a imagem Docker
```bash
docker build -t go-alimentos .
```

 ### Execute o contêiner Docker
```bash
docker run -p 8080:8080 go-alimentos
```

O aplicativo estará disponível em http://localhost:8080.

## Exemplo de Uso

### 1. Obter todos os alimentos

```bash
`curl http://localhost:8080/tabela-alimentos` 
```
### 2. Buscar alimentos por nome

```bash
`curl http://localhost:8080/buscar-alimento?nome=Arroz` 
```
## Contribuições

Sinta-se à vontade para contribuir, abrir problemas ou propor melhorias para este projeto.

## Licença

Este projeto está licenciado sob a Licença MIT - consulte o arquivo [LICENSE.md](#) para obter detalhes.

