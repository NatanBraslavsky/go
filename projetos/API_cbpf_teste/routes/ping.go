package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}