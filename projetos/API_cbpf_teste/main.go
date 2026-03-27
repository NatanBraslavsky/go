package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"exemplo.com/api-cbpf/env"
	"exemplo.com/api-cbpf/mongo"
	"exemplo.com/api-cbpf/routes"
)

func main() {
	_ = godotenv.Load("./.env")
	env.Load()
	mongo.InitDB()
	mongo.Setup()

	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic("Failed to start gin: " + err.Error())
	}
}
