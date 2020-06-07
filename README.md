# Starwars RESTAPI 


## Getting Started

Esta API REST fornece dados dos planetas presentes na franquia StarWars.

* Para cada planeta, os seguintes dados são obtidos do banco de dados da aplicação (MongoDB), sendo inserido manualmente:

* Nome
* Clima
* Terreno
 
 ```
 NOTA: Ao realizar uma inserção manual, a quantidade de aparições em filmes será registrada para o filme inserido.
 Filmes não presentes na franquia Starwars não serão registrados. 
 ``` 

* Para cada planeta também é retornarnada a quantidade de aparições em filmes, as quais são obtidas pela API pública do Star Wars: https://swapi.dev/about


### Prerequisitos

```
Mux [go get -u github.com/gorilla/mux]
MongoDB Local [https://www.mongodb.com/download-center/community?tck=docs_server]
```

### Funcionalidades

As seguintes funcionalidades são disponibilizadas:

```
* Adicionar um planeta (com nome, clima e terreno)
* Listar planetas
* Buscar por nome
* Buscar por ID*
* Remover planeta

```

## Running the tests

Testes unitários automáticos podem ser executados a partir do terminal com o comando abaixo (para o main_test.go)

```
go test
```

Para execução de testes a partir de API Clients (Ex.: Postman), os exemplos abaixo são disponibilizados:

### Adicionar um planeta (com nome, clima e terreno)
```
[POST] http://localhost:8000/api/planets
{
    "name": "Alderaan",
    "climate": "tropical",
    "terrain": "jungle rainforests"
}
```

### Listar planetas
```
[GET] http://localhost:8000/api/planets

```
### Buscar por nome
```
[GET] http://localhost:8000/api/planets/search/Alderaan

```
### Buscar por ID*
```
[GET] http://localhost:8000/api/planets/5edd519fab0b1b28b274f6d0
```
### Remover planeta
```
[DELETE] http://localhost:8000/api/planets/5edd519fab0b1b28b274f6d0
```


## Built With

* [Golang](http://golang.org/) - go1.12.4 windows/amd64
* [MongoDB](www.mongodb.com) - v4.2.7

