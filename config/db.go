package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	SQLDB  *sql.DB  // pure SQL ke liye
	GormDB *gorm.DB // ORM ke liye
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// 1. GORM Connection
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect using GORM:", err)
	}
	GormDB = gormDB

	// 2. SQL Connection (for manual query)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect using SQL:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("SQL Database not reachable:", err)
	}
	SQLDB = sqlDB
}
