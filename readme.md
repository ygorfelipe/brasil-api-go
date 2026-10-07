## Estudo API GO LANG - Brasil API

' Para iniciar um projeto sempre devemos criar um repositorio git 
comando:
    go mod init 'github.com/ygorfelipe/brasil-api-go' // deve ser sem o HTTPS

dessa forma é possível reutilizar o esse modulo, basicamente é um pub.dev
ele cria um go.mod (pubspec.yaml) - sempre tem retro-compatibilidade do go

* Projeto brasil API (via cep) - GO com Gin Gonic

estrutura de pastas


│   go.mod
│   go.sum
│   readme.md
│   
├───cmd
│       main.go // a main sempre ira ter o package main, nunca o padrão cmd     
│       
└───internal // tudo onde a nossa aplicação irá ficar e não compartilhada externamente.
    ├───app
    │   └───bootstrap 
    │           server.go // sera inicializada o server
    ├───domain
    └───infra

GO não é POO - porém é possível trabalhar como se fosse
funções camelCase é compartilhada externamente caso ao contrario é privada, só sera acessada dentro do próprio pacote