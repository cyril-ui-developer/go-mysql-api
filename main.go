package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type Employee struct {
	ID        int
	FirstName string `json: "firstName`
	LastName string `json: "lastName"`
	Address   string `json: "address"`
}

func jsonRReesponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonRReesponse(w, map[string]string{"error": err.Error()})
		return
	}
	w.Write(b)
}

func main() {
	// Get the port from env vars or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting the server up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go and MySQL API"))
	})
	r.Get("/employees", func(w http.ResponseWriter, r *http.Request) {
		employees := []Employee{
			Employee{
				LastName: "Cyril", FirstName: "Ajieh",
			},
			Employee{
				LastName: "Campbell", FirstName: "David",
			},
			Employee{
				LastName: "Doe", FirstName: "John",
			},
		}
		// jsonRReesponse(w, map[string]string{"name": "Cyril"})
		jsonRReesponse(w, employees)
	})
	fmt.Println("Welcome to Go and MySQL API")

	log.Fatal(http.ListenAndServe(":"+port, r))
}

