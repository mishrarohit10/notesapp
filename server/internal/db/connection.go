package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pining database: %v", err)
		return err
	}

	log.Println("Database connection succesfull")
	return nil
}
