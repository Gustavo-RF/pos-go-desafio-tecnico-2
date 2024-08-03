# Desafio técnico 2
Esse desafio tem como objetivo criar um stress test, recebendo uma url de teste e a quantidade de concorrencias e chamadas, e executar as requisições para a url.

## Como funciona
O programa receberá do usuário três parâmetros:
- url: Url que será chamado no teste
- concurrency: total de workers que trabalharão simultâneamente
- requests: total de requests a serem feitas

Após execução das requests, será mostrado um resumo dos resultados no seguinte formato:
```
Stress test finalizado em xxxxxs
Total de requests realizadas:           xxx
Status xxx                      Qtd:    xx
Status xxx                      Qtd:    xx
...
```

## Executando o projeto
Primeiramente é necessário fazer o build do projeto com o docker. Vá até a pasta do projeto e digite no terminal:
```
docker build -t desafio-tecnico-2 .
```

Isso gera a imagem do projeto no docker. Para fazer os testes passando os parâmetros, execute:
```
docker run --rm --name desafio-tecnico-2 desafio-tecnico-2 --url=your_url --concurrency=qtd_concurrents --requests=qtd_requests
```
> Caso utilize outra tag no build, ela deverá ser usada da mesma forma no run

> o parâmetro --rm serve para matar o container logo após a execução e não precisar ser apagado manualmente, facilitando os testes

## Exemplos
- Podemos utilizar o resuldado do desafio técnico 1 para testar:
  ```
  docker run --rm --name desafio-tecnico-2 desafio-tecnico-2 --url=http://ip_do_seu_docker:8080 --concurrency=10 --requests=100
  ```
  > É necessário passar o ip do docker caso esteja subindo pelo docker
  
  ```
  Stress test finalizado em 161.992722ms
  Total de requests realizadas:           100
  Status 200                      Qtd:    5
  Status 429                      Qtd:    95
  ```

- Pode ser testado também utilizando o [httpstatus](https://httpstat.us/), um serviço de teste que retorna randomicamente um status de resposta passando o range de status
  ```
  docker run --rm --name desafio-tecnico-2 desafio-tecnico-2 --url=http://httpstat.us/random/200-209 --concurrency=10 --requests=100
  ```
  > chamando a url ```http://httpstat.us/random/200-209``` retornará um status aleatório entre 200 e 209
  
  ```
  Stress test finalizado em 32.704407004s
  Total de requests realizadas:           100
  Status 200                      Qtd:    9
  Status 201                      Qtd:    14
  Status 202                      Qtd:    6
  Status 203                      Qtd:    5
  Status 204                      Qtd:    10
  Status 205                      Qtd:    11
  Status 206                      Qtd:    14
  Status 207                      Qtd:    13
  Status 208                      Qtd:    11
  Status 209                      Qtd:    7
  ```

