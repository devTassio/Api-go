package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()  // Create Gin engine
	initializeRoutes(router) // Configure routes

	router.Run() // Start the server
}
