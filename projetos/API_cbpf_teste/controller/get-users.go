package controller

import (
	"github.com/gin-gonic/gin"

	"exemplo.com/api-cbpf/user"
)

func GetUsers(c *gin.Context) {
	users, err := user.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, users)
}
