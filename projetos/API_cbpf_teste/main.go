package main

import (
	"exemplo.com/api-cbpf/env"
	"exemplo.com/api-cbpf/mongo"
	"exemplo.com/api-cbpf/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("./.env")
	env.Load()
	mongo.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}