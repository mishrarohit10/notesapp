package main

import (
	"fmt"
	"notesapp/server/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi from the server")

	router := gin.Default()

	setupRoutes(router)

	if err := router.Run(); err != nil {
		panic(err)
	}
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
