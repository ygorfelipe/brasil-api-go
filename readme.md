## Estudo API GO LANG - Brasil API

' Para iniciar um projeto sempre devemos criar um repositorio git 
comando:
    go mod init 'github.com/ygorfelipe/brasil-api-go' // deve ser sem o HTTPS

dessa forma é possível reutilizar o esse modulo, basicamente é um pub.dev
ele cria um go.mod (pubspec.yaml) - sempre tem retro-compatibilidade do go

* Projeto brasil API (via cep) - GO com Gin Gonic

estrutura de pastas



brasil-api-go
│   go.mod
│   go.sum
│   readme.md
│   
├───cmd
│       main.go // a main sempre ira ter o package main, nunca o padrão cmd     
│       
└───internal // tudo onde a nossa aplicação irá ficar e não compartilhada externamente.
    ├───app
    │   ├───bootstrap
    │   │       server.go // sera inicializada o server
    │   │       
    │   └───handlers
    │       └───locations
    │           │   location_handler.go
    │           │   
    │           └───dto
    │                   state_response.go
    │                   
    ├───domain
    │   └───entities
    │           states_entity.go
    │           
    └───infra
        └───repositories
            └───location
                │   location_repository.go
                │   
                └───dto
                        brasil_api_response.go




Forma da qual irei trabalhar e a compreensão com base no meu conhecimento.

A estrutura abaixo é a forma que acho mais facil de se trabalhar


Primeiras camadas para ser criadas

1° Criação da seguinte estrutura de pastas

├───cmd
│       main.go     
│       
└───internal
    ├───app
    │   ├───bootstrap
    │   │       server.go


Configuração principal esta na main.go
Iremos criar uma função StartServer, onde sera configurado a API, ela foi feita inicialmente com gin-gonic, podendo ser refeita para 100% go

Foi criado outra função chamado configureRoutes, com passagem de paramentros, no qual ele esta pegando apenas a instancia de memoria Engine (*) do gin-gonic.

Dentro da função configureRoutes, foi realizado a injeção de dependência de forma manual (é possível realizar essa injeção via pacote)

Após a configuração inicial da API em server.go, o fluxo da qual achei ideal trabalhar é da seguinte forma.

2º Trabalhando por camadas

├───domain
    │   └───entities
    │           states_entity.go

Criando a class de modelo, após a criação da classe de modelo, proximo passo é a criação da camada de infra/repositories

└───infra
    └───repositories
        └───location
            │   location_repository.go
            │   
            └───dto
                    brasil_api_response.go


Primeira a se fazer é realizar o decoded (json_serializable), ou seja, a criação da dto dentro de infra, pelo fato de que logo após sera utilizada ela em código.
No camada de repository, é a camada da igual irá buscar os dados do end-point




