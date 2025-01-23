package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Microservice running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
