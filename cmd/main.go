package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ivar1309/Api-Go-Boilerplate/config"
	"github.com/ivar1309/Api-Go-Boilerplate/internal/routes"
)

func main() {
	config.LoadEnv() // Load environment variables

	router := routes.SetupRouter() // Set up API routes

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
