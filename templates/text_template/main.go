package main

import (
	"os"
	"text/template"
)

type Person struct {
	Fname string
	Lname string
}

func main() {
	// Define the template with placeholders
	tmpl, err := template.New("greeting").Parse("Hello, {{.Fname}} {{.Lname}}!")
	if err != nil {
		panic(err)
	}

	// Provide data for the template
	data := Person{Fname: "John", Lname: "Doe"}

	// Execute the template and write output to the console
	err = tmpl.Execute(os.Stdout, data)
	// output - Hello, John Doe!

	if err != nil {
		panic(err)
	}
}
