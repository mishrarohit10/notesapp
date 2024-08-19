package user

import (
	"net/http"
	"notesapp/server/internal/db"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	type Login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var creds Login
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// check if email is present in db or not
	var exists bool
	err := db.DB.Debug().Raw("SELECT EXISTS (SELECT 1 FROM users WHERE email=?)", creds.Email).Scan(&exists).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This email has no account"})
		return
	}

	var storedPassword string
	err = db.DB.Debug().Raw("SELECT password FROM users WHERE email=?", creds.Email).Scan(&storedPassword).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password not found"})
		return
	}

	// check if password is correct or not
	if creds.Password != storedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged In", "error": false})
}

func SignUp(c *gin.Context) {

	type SignUp struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var creds SignUp
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error bindingJSON", "error": true})
		return
	}

	// check if email is present in db or not
	var exists bool
	err := db.DB.Debug().Raw("SELECT EXISTS (SELECT 1 FROM users WHERE email=?)", creds.Email).Scan(&exists).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying email database"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
		return
	}

	err = db.DB.Debug().Exec("INSERT INTO users (email, password) VALUES (?,?)", creds.Email, creds.Password).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signed Up", "email": creds.Email, "error": false})
}
