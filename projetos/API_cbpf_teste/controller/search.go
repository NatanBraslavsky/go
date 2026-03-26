package controller

import (
	"github.com/gin-gonic/gin"
)

func GetSearch(c *gin.Context) {
	query := c.Query("q")
	c.JSON(200, gin.H{"q": query})
}
