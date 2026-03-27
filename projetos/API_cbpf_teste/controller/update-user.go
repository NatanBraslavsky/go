package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/mongo"
	"exemplo.com/api-cbpf/user"
	"exemplo.com/api-cbpf/utils"
)

func UpdateUser(c *gin.Context) {
	objID, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})

		return
	}

	var updatedUser mongo.User

	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data"})

		return
	}

	if !utils.IsValidEmail(updatedUser.Email) {
		c.JSON(400, gin.H{"message": "Invalid email"})

		return
	}

	err = user.UpdateByID(c.Request.Context(), objID, updatedUser)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(404, gin.H{"message": "User not found"})

			return
		}

		c.JSON(500, gin.H{"message": "Error updating user"})

		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}
