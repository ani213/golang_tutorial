package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Person struct represents an individual with a first and last name
type Person struct {
	FirstName string
	LastName  string
}

// removeSpaces removes all spaces from a string
func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func main() {
	// Template file location
	templateFile := "./personlist.tmpl"

	// Define a function map for use in templates
	fmap := template.FuncMap{
		"removeSpaces": removeSpaces,    // Custom function to remove spaces
		"capitalize":   strings.ToUpper, // Built-in function to convert to uppercase
	}

	// Define handler for "/shortlist/"
	http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the template with the function map
		tp, err := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(templateFile)
		if err != nil {
			http.Error(w, "Template parsing error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		// Sample list of people
		people := []Person{
			{"Bo  b", "Bark   er"},
			{"Lar ry", "F  lint"},
		}

		// Execute template
		err = tp.Execute(w, people)
		if err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	// Define handler for "/longlist/"
	http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
		tp, err := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(templateFile)
		if err != nil {
			http.Error(w, "Template parsing error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		people := []Person{
			{"Bob", "Barker"},
			{"Lar  ry", "Flint"},
			{"Su san", "Sa ra ndo n"},
			{"Br ad", "Pit t"},
			{"Bet ty", "Whi te"},
		}

		err = tp.Execute(w, people)
		if err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	// Define handler for "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tp, err := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(templateFile)
		if err != nil {
			http.Error(w, "Template parsing error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		err = tp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	// Start the HTTP server on port 8080
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server startup error:", err)
	}
}
