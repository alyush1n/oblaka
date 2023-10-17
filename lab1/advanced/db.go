package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func InitDB() error {

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := "disable"
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s", dbUser, dbPassword, dbName, sslMode, dbHost, dbPort)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	err = CreateMessagesTable()
	if err != nil {
		return err
	}

	return nil
}

func CreateMessagesTable() error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message VARCHAR(255)
)
`)
	return err
}

func InsertMessage(message string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO messages (message) VALUES ($1) RETURNING id", message).Scan(&id)
	return id, err
}
