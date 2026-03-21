package routes

import (
	"exemplo.com/rest-api/middewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middewares.Authenticate)
	authenticated.POST("/events", middewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
