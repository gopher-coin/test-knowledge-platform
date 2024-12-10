package main

import (
	"fmt"
	"log"
	"net/http"
	"test-knowledge-platform/internal/repositories"
	"test-knowledge-platform/internal/routes"
)

func main() {

	err := repositories.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer repositories.CloseDB()

	routes.SetupRotes()

	fmt.Println("Server is starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
