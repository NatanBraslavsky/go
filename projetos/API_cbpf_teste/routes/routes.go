package routes

import (
	"exemplo.com/api-cbpf/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/ping", controller.GetPing)
	server.GET("/users/:id", controller.GetUser)
	server.GET("/search", controller.GetSearch)

	server.POST("/users", controller.CreateUser)
}