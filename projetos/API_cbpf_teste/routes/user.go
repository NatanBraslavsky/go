package routes

import (
	"net/http"

	"exemplo.com/api-cbpf/models"
	"github.com/gin-gonic/gin"
)

func getUser(context *gin.Context) {
	id := context.Param("id")

	context.JSON(http.StatusOK, id)
}

func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusOK, user)
}