package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Load and parse index.html
	tmpl := template.Must(template.ParseFiles("index.html"))

	// template.ParseFiles("index.html")-
	// Loads the "index.html" file from disk.
	// Parses it into a *template.Template object.
	// Returns this object (or an error if parsing fails).

	// template.Must(...)-
	// Wraps template.ParseFiles(...) to handle errors automatically.
	// If parsing fails, Must panics instead of returning an error.

	// Data to pass to the template
	data := PageData{
		Title:   "Welcome Page",
		Message: "Hello, Go Templates!",
	}

	// Execute template with data
	tmpl.Execute(w, data)
	// output -
	// Welcome Page (title)
	// Hello, Go Templates!

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
