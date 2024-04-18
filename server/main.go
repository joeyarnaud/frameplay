package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {
	// Simple handler for the path
	http.HandleFunc("/", handlers.HandleRoot)
	println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
