package routes

import (
	"net/http"
	"test-knowledge-platform/internal/handlers"
)

func SetupRotes() {
	http.HandleFunc("/api/users", handlers.CreateUserHandler)
}
