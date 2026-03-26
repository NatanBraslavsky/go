package routes

import (
	"exemplo.com/rest-api/controller"
	"exemplo.com/rest-api/middewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controller.GetEvents)
	server.GET("/events/:id", controller.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middewares.Authenticate)
	authenticated.POST("/events", middewares.Authenticate, controller.CreateEvent)
	authenticated.PUT("/events/:id", controller.UpdateEvent)
	authenticated.DELETE("events/:id", controller.DeleteEvent)
	authenticated.POST("events/:id/register", controller.RegisterForEvent)
	authenticated.DELETE("events/:id/register", controller.CancelRegistration)

	server.POST("/signup", controller.Signup)
	server.POST("/login", controller.Login)
}
