package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/user"
	"exemplo.com/api-cbpf/utils"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data"})

		return
	}

	if !utils.IsValidEmail(newUser.Email) {
		c.JSON(400, gin.H{"message": "Invalid email"})

		return
	}

	newUser.ID = bson.NewObjectID()

	err = user.Insert(c.Request.Context(), newUser)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error saving user"})

		return
	}

	c.JSON(200, newUser)
}
