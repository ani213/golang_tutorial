package main

import (
	"fmt"
	"net/http"
)

// PersonHandler struct to hold personal details
type PersonHandler struct {
	FirstName string
	LastName  string
}

// CarHandler struct to hold car details
type CarHandler struct {
	Make  string
	Model string
	Year  int
}

// ServeHTTP method for PersonHandler to handle HTTP requests
func (h *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name: %v %v", h.FirstName, h.LastName)
}

// ServeHTTP method for CarHandler to handle HTTP requests
func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", h.Make, h.Model, h.Year)
}

// creating handle function for person
func handlePerson(w http.ResponseWriter, r *http.Request) {
	firstName := "Bob"
	lastName := "Smith"
	fmt.Fprintf(w, "Name: %v %v", firstName, lastName)
}

// creating handle function for car
func handleCar(w http.ResponseWriter, r *http.Request) {
	make := "Ford"
	model := "Fiesta"
	year := 2005
	fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", make, model, year)
}

func main() {
	// Registering handlers for different routes
	http.Handle("/person", &PersonHandler{"Bob", "Smith"})
	http.Handle("/car", &CarHandler{"Ford", "Fiesta", 2005})

	//using handler function
	http.HandleFunc("/handleperson", handlePerson)
	http.HandleFunc("/handlecar", handleCar)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
