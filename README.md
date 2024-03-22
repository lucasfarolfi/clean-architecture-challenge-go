# clean-architecture-challenge-go


## Como rodar?

1. utilize o comando "docker-compose up --build" para executar o docker para criar o banco de dados e o RabbitMQ

2. utilize o comando "go mod tidy" para fazer a instalacao dos pacotes necessarios para executar o projeto

3. Apos a execucao correta do docker-compose, rode o comando "cd cmd/orderystem" e em seguida execute "go run main.go wire_gen.go" para executar o projeto

Obs1: Este servidor estara disponivel na porta 8080
porta 8000, o gRPC server estara na porta 50051 e o GraphQL estara na porta 8080.

Obs2: Uma migration para a criacao da tabela de orders no banco sera executada apos a execucao do app

## Como executar o use case de ListOrders

### WebServer
Este servidor estara disponivel na porta 8000

Para utilizar execute uma requisicao "GET http://localhost:8000/orders", ou rode o arquivo "api/list_orders.http" 

### gRPC
Este servidor estara disponivel na porta 50051

Para utilizar execute os seguintes passos:
1. rode "evans -r repl" no terminal
2. rode o comando "package pb"
3. rode o comando "call ListOrders"

### GraphQL
Este servidor estara disponivel na porta 8080

Para utilizar, cole a seguinte URL no navegador "http://localhost:8080" 

Apos isso, cole a seguinte query na interface e execute

```
query queryOrders{
  orders{
    id
    Price
    Tax
    FinalPrice
  }
}
```