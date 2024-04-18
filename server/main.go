package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to parse the JSON input.
type Person struct {
	Email string `json:"email"`
	Password  string    `json:"password"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w, r)

	// Check if it's a preflight request and handle accordingly
	if r.Method == "OPTIONS" {
		return // preflight request doesn't need further handling
	}

	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON body
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate the input
	if person.Email == "" {
		http.Error(w, "Invalid input: Email is empty", http.StatusBadRequest)
		return
	}

	// Log the input to STDOUT
	log.Printf("Received: %+v\n", person)

	// Respond to the client
	fmt.Fprintf(w, "Received: %+v", person)

	// Optionally, send the input to another HTTP API
	forwardInput(person)
}

func forwardInput(person Person) {
	// Here we serialize the person struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// Sending a POST request to another API
	resp, err := http.Post("http://example.com/api/target", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending request to API: %v", err)
	}
	defer resp.Body.Close()

	// Optionally, log the response from the API
	log.Println("Response Status:", resp.Status)
}

func setCORSHeaders(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "http://localhost:3000" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

func main() {
	http.HandleFunc("/", handler)
	println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
