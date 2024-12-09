package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "A project for testing students' knowledge")
	})
	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
}
