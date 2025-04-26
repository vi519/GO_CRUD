package main

import (
	"log"
	"net/http"

	"go-crud-live/config"
	"go-crud-live/routes"
)

func main() {
	config.Connect()

	router := routes.RegisterRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
