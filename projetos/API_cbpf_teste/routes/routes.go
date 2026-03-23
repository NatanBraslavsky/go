package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/ping", getPing)
	server.GET("/users/:id", getUser)
	server.GET("/search", getSearch)

	server.POST("/users", createUser)
}