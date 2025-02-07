package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {

	templateFile := "./personlist.tmpl"

	http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)

		people := []Person{
			{"Bob", "Barker"},
			{"Larry", "Flint"},
			{"Susan", "Sarandon"},
			{"Brad", "Pitt"},
			{"Betty", "White"}}

		tp.Execute(w, people)

	})

	http.ListenAndServe(":8080", nil)
}
