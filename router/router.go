package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()  // Criar o mecanismo Gin
	initializeRoutes(router) // Configurar rotas

	router.Run() // Iniciar o servidor
}
