package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string   `json:"message"`
	Version  string `json:"version"`
	Routes  []string `json:"routes"`
}

func main() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			response := Response{
				Message: "Simple Golang RESTful API example",
				Version: "0.1.0-beta",
				Routes:  []string{"/api/name"},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	})

	http.HandleFunc("/api/name", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			name := r.URL.Query().Get("name")
			if name == "" {
				http.Error(w, "Name parameter is missing", http.StatusBadRequest)
				return
			}
			fmt.Fprintf(w, "Hello, %s!", name)
		}
	})

	http.ListenAndServe(":8080", nil)
}
