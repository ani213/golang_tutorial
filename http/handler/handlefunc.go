package main

import (
	"fmt"
	"net/http"
)

// Handler function for /person route with dynamic query parameters
func personHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters from URL
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")

	// Provide default values if parameters are missing
	if firstName == "" {
		firstName = "John"
	}
	if lastName == "" {
		lastName = "Doe"
	}

	fmt.Fprintf(w, "Name: %v %v", firstName, lastName)
}

// Handler function for /car route with dynamic query parameters
func carHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters from URL
	make := r.URL.Query().Get("make")
	model := r.URL.Query().Get("model")
	year := r.URL.Query().Get("year")

	// Provide default values if parameters are missing
	if make == "" {
		make = "Toyota"
	}
	if model == "" {
		model = "Corolla"
	}
	if year == "" {
		year = "2020"
	}

	fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", make, model, year)
}

func mains() {
	// Registering handler functions
	http.HandleFunc("/person", personHandler)
	http.HandleFunc("/car", carHandler)

	// Start the HTTP server
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
