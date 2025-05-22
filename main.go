package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlePort8080(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from port 8080!")
}

func handlePort8081(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from port 8081!")
}

func main() {
	// Create separate servers for each port
	server8080 := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handlePort8080),
	}

	server8081 := &http.Server{
		Addr:    ":8081",
		Handler: http.HandlerFunc(handlePort8081),
	}

	// Start server on port 8080 in a goroutine
	go func() {
		log.Printf("Server starting on port 8080...")
		if err := server8080.ListenAndServe(); err != nil {
			log.Printf("Error starting server on port 8080: %v", err)
		}
	}()

	// Start server on port 8081
	log.Printf("Server starting on port 8081...")
	if err := server8081.ListenAndServe(); err != nil {
		log.Printf("Error starting server on port 8081: %v", err)
	}
}
