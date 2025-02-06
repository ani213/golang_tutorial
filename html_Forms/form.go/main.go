package main

import (
	"fmt"
	"net/http"
)

// Function to render the form page
func formHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl, _ := template.ParseFiles("form.html")
	// tmpl.Execute(w, nil)
	// or
	http.ServeFile(w, r, "form.html")
}

// Function to handle form submission
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Retrieve form values
	name := r.FormValue("name")
	email := r.FormValue("email")

	// Respond to the user
	fmt.Fprintf(w, "Received:\nName: %s\nEmail: %s", name, email)
}

func main() {
	// Serve static HTML form
	http.HandleFunc("/", formHandler)

	// Handle form submission
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
