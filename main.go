package main

import (
	"fmt"
	config "go-crud-mvc/db"
)

func main() {
	// Call your connection_db function here
	config.Connection_db()

	// You can now use the DB connection, perform queries, etc.
	fmt.Println("Main function is running")
}
