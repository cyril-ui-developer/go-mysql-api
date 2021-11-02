package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	// Get the port from env vars or default to 8000
	port := "8000"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting the server up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go and MySQL API"))
	})
	fmt.Println("Welcome to Go and MySQL API")

	log.Fatal(http.ListenAndServe(":"+port, r))
}