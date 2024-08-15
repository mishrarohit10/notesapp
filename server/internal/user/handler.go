package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	type Login struct {
		username string
		password string
	}

	var creds Login
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged In", "email": creds.username, "password": creds.password})
}
