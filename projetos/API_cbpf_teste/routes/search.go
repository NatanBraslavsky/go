package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSearch(context *gin.Context) {
	query := context.Query("q")
	context.JSON(http.StatusOK, gin.H{"q": query})
}