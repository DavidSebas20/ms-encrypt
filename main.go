package main

import (
	"log"
	"net/http"
)

// Hash struct
type HashRequest struct {
	Password string `json:"password"`
}

type HashResponse struct {
	Hash string `json:"hash"`
}

// Verify struct
type VerifyRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

type VerifyResponse struct {
	Match bool `json:"valid"`
}

func main() {
	log.Println("Microservice running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
