package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type ErrorResponse struct {
    Error   bool   `json:"error"`
    Message string `json:"message"`
}

type SuccessResponse struct {
    Message string `json:"message"`
}

func SetCORSHeaders(w http.ResponseWriter, r *http.Request) {
	// allows the react client 
	origin := r.Header.Get("Origin")
	if origin == "http://localhost:3000" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

// create simple handler for error messages
func HandleError(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := ErrorResponse{
        Error:   true,
        Message: message,
    }
    json.NewEncoder(w).Encode(response)
}

// create simple handler for success messages
func HandleSuccess(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := SuccessResponse{
        Message: message,
    }
    json.NewEncoder(w).Encode(response)
}

// for localhost:8080/
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// so we can do locally 
	SetCORSHeaders(w, r)

	if r.Method == "OPTIONS" {
		return 
	}

	// currently only post is allowed
	if r.Method == "POST" {
		HandleSignIn(w, r)
	}

	if r.Method != "POST" {
		HandleError(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
}

func ForwardInput(person Person) {
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// Sending a POST to another api input your target here 
	resp, err := http.Post("http://example.com/api/target", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending request to API: %v", err)
	}
	defer resp.Body.Close()

	log.Println("Forward API Response Status:", resp.Status)
}