package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error pining database: %v", err)
		return err
	}

	log.Println("Database connection succesfull")
	return nil
}
