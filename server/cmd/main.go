package main

import (
	"log"
	"net/http"
	"notesapp/server/internal/db"
	"notesapp/server/internal/user"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	log.Println("Current working directory ->", cwd)

	envPath := filepath.Join(cwd, "../../server/.env")

	log.Println("Environment variable path ->",envPath)

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf(".env file does not exist at path: %s", envPath)
	}

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}

	err = db.Init()
	if err != nil {
		log.Fatalf("Failed to connect to Database")
	}

	log.Println("Database connect succesfull")

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
		userGroup.POST("/signup", user.SignUp)
	}
}
