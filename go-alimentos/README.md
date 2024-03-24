# Tabela de Alimentos API em Golang

Esta é uma API simples em Golang para acessar uma tabela de alimentos. A tabela é carregada a partir de um arquivo JSON local e fornece endpoints para obter todos os alimentos ou buscar alimentos por nome.

## Pré-requisitos

Certifique-se de ter o Go instalado em seu ambiente de desenvolvimento.

## Telas

![Screenshot 2024-03-24 at 13-58-18 Screenshot](https://github.com/marcelosanto/tabela_taco/assets/11478538/df2e96e5-fd42-4beb-9e36-909f18e40786)

![Screenshot 2024-03-24 at 13-58-47 Screenshot](https://github.com/marcelosanto/tabela_taco/assets/11478538/39be498d-4b7d-4cd3-8f3a-f5617daf2b48)


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

O servidor será iniciado em http://localhost:5000.

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
docker run -p 5000:5000 go-alimentos
```

O aplicativo estará disponível em http://localhost:5000.

## Exemplo de Uso

### 1. Obter todos os alimentos

```bash
`curl http://localhost:5000/tabela-alimentos` 
```
### 2. Buscar alimentos por nome

```bash
`curl http://localhost:5000/buscar-alimento?nome=Arroz` 
```
## Contribuições

Sinta-se à vontade para contribuir, abrir problemas ou propor melhorias para este projeto.

## Licença

Este projeto está licenciado sob a Licença MIT - consulte o arquivo [LICENSE.md](#) para obter detalhes.

