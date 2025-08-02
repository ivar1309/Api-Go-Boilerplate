package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ivar1309/Api-Go-Boilerplate/internal/db"
	"github.com/ivar1309/Api-Go-Boilerplate/internal/routes"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	router := routes.SetupRouter() // Set up API routes

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
