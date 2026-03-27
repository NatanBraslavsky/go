package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/user"
)

func GetUser(c *gin.Context) {
	objID, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})

		return
	}

	foundUser, err := user.GetByID(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			c.JSON(404, gin.H{"message": "User not found"})

			return
		}

		c.JSON(500, gin.H{"message": "Error fetching user"})

		return
	}

	c.JSON(200, foundUser)
}
