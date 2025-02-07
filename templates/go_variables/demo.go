package main

import (
	"html/template"
	"log"
	"net/http"
)

// Person struct represents an individual with a first and last name
type Person struct {
	FirstName string
	LastName  string
}

func main() {
	// Define the template file location
	templateFile := "./personlist.tmpl"

	// Handler for "/shortlist/"
	http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {
		tp, err := template.ParseFiles(templateFile)
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		// Short list of people
		people := []Person{
			{"Bob", "Barker"},
			{"Larry", "Flint"},
		}

		// Execute the template
		err = tp.Execute(w, people)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	// Handler for "/longlist/"
	http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
		tp, err := template.ParseFiles(templateFile)
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		// Long list of people
		people := []Person{
			{"Bob", "Barker"},
			{"Larry", "Flint"},
			{"Susan", "Sarandon"},
			{"Brad", "Pitt"},
			{"Betty", "White"},
		}

		// Execute the template
		err = tp.Execute(w, people)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	// Handler for "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tp, err := template.ParseFiles(templateFile)
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		// Execute the template with no data
		err = tp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	// Start the HTTP server
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server startup error:", err)
	}
}
