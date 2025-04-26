package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the DB connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Open the DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) // Log the error and terminate the application
	}

	// Check if the connection is established successfully
	if err = db.Ping(); err != nil {
		log.Fatal(err) // Log the error and terminate the application
	}

	// Set the global DB variable
	DB = db

	// Ensure DB connection is closed gracefully when the app exits
	// So that connections are properly cleaned up when app terminates
	defer func() {
		if err := DB.Close(); err != nil {
			log.Println("Error closing DB connection:", err)
		}
	}()
}
