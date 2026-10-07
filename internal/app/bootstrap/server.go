package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// para chamar a função, em outro pacote, basta informar o nome do pacote e sua função ex: bootstrap.SrtartServer

func StartServer() {
	// Declarando as configurações e inicializando o server
	// por convenção as var são declaradas como antigamente ex: engine = e; webhook = wb;

	// AQUI é onde irá ficar todas as configurações e chamadas do server.
	// geralmente em GO é criado sempre em um unico arquivo varias funcoes de chamadas de configuracoes
	// APOS DEFAULT E ONDE IREMOS COLOCAR AS CONFIGURACOES
	e := gin.Default()
	configureRoutes(e)

	// em go não tem tryCatch, sempre sera Either e Nil, podendo ser multiplos retornos
	err := e.Run(":8080")

	if err != nil {
		panic(err) // panic para o projeto e print o erro
	}

	fmt.Println("Server Started")
}

// função para criar rotas
// quando utiliza o "*" você trabalha com PONTEIROS DE MEMORIA (PROCURAR SOBRE)
// no qual você pega apenas a referência da funçao do pacote que você deseja, ou seja, não precisa extrair o pacote todo
// para depois escolher o que trabalhar, basta escolher diretamente o que precisa em dart ex: as env, as pkg; ao pado do import
func configureRoutes(e *gin.Engine) {
	// G grupo de rotas

	g := e.Group("/api/v1")
	{
		// em go nós temos sempre o nosso contexto quando precisaremos
		g.GET("/states", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

}
