package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func Connection_db() {
	var err error
	connStr := "host=localhost port=5432 user=vineet password=batman@123 dbname=postgres sslmode=disable"

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database is connected")

}
