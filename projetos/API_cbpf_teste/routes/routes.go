package routes

import (
	"github.com/gin-gonic/gin"

	"exemplo.com/api-cbpf/controller"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/ping", controller.GetPing)

	server.POST("/users", controller.CreateUser)
	server.GET("/users/:id", controller.GetUser)
	server.PUT("users/:id", controller.UpdateUser)
	server.DELETE("users/:id", controller.DeleteUser)
	server.GET("/users", controller.GetUsers)
}
