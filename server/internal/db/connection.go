package db

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func Init() error {
	
	connStr := fmt.Sprintf("host=localhost port=5432 user=postgres password=%s dbname=postgres sslmode=disable", os.Getenv("POSTGRES_PASSWORD"))

	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	return DB.Ping()
}
