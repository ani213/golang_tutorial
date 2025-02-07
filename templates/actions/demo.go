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

	http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)

		people := []Person{
			{"Bob", "Barker"},
			{"Larry", "Flint"}}

		tp.Execute(w, people)

	})

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)

		tp.Execute(w, nil)

	})

	http.ListenAndServe(":8080", nil)
}
