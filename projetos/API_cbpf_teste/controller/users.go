package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/mongo"
	"exemplo.com/api-cbpf/utils"
)

func GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	var users []models.User

	cursor, err := mongo.ApiProject.Collection("user").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})

			return
		}
	}()

	for cursor.Next(ctx) {
		var user models.User

		err := cursor.Decode(&user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})

			return
		}

		users = append(users, user)
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})

		return
	}

	var user models.User

	err = mongo.ApiProject.Collection("user").FindOne(c, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})

		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data"})

		return
	}

	// valida email
	if !utils.IsValidEmail(user.Email) {
		c.JSON(400, gin.H{"message": "Invalid email"})

		return
	}

	user.ID = bson.NewObjectID()

	_, err = mongo.ApiProject.Collection("user").InsertOne(c, user)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error saving user"})

		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})

		return
	}

	result, err := mongo.ApiProject.Collection("user").DeleteOne(c, bson.M{"_id": objID})
	if err != nil {
		c.JSON(500, gin.H{"message": "Error deleting user"})

		return
	}

	if result.DeletedCount == 0 {
		c.JSON(404, gin.H{"message": "User not found"})

		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})

		return
	}

	var updatedUser models.User

	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data"})

		return
	}

	// valida email
	if !utils.IsValidEmail(updatedUser.Email) {
		c.JSON(400, gin.H{"message": "Invalid email"})

		return
	}

	result, err := mongo.ApiProject.Collection("user").UpdateOne(
		c,
		bson.M{"_id": objID},
		bson.M{"$set": updatedUser},
	)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error updating user"})

		return
	}

	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"message": "User not found"})

		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}
