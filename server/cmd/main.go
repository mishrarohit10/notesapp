package main

import (
	"fmt"
	"net/http"
	"notesapp/server/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi from the server")

	router := gin.Default()

	router.GET("/", rootHandler)

	setupRoutes(router)

	if err := router.Run(); err != nil {
		panic(err)
	}
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "working fine"})
}

func setupRoutes(router *gin.Engine) {
	userRouter(router)
}

func userRouter(router *gin.Engine) {
	userGroup := router.Group("/")
	{
		userGroup.POST("/login", user.Login)
	}
}
