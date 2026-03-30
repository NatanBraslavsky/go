package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/user"
)

type DeleteUserInput struct {
	Password string `binding:"required" json:"password"`
}

func DeleteUser(c *gin.Context) {
	objID, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})

		return
	}

	var input DeleteUserInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"message": "Cold not parse request data"})

		return
	}

	foundUser, err := user.GetByID(c.Request.Context(), objID)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})

		return
	}

	if foundUser.Password != input.Password {
		c.JSON(401, gin.H{"message": "Invalid password"})

		return
	}

	err = user.DeleteByID(c.Request.Context(), objID)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error deleting user"})

		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
