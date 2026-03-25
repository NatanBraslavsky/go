package controller

import (
	"exemplo.com/api-cbpf/models"
	"exemplo.com/api-cbpf/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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

	c.JSON(200, id)
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	result, err := mongo.ApiProject.Collection("user").InsertOne(c, user)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error saving user"})
		return
	}

	user.ID = result.InsertedID.(bson.ObjectID)

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
