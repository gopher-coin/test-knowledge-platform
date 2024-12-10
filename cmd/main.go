package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is starting on http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Test Knowledge Platform!")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
