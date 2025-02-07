package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	// Define a template function map (FuncMap) that includes a custom function.
	// "toUpper" is mapped to strings.ToUpper, which converts a string to uppercase.
	funcs := template.FuncMap{
		"toUpper": strings.ToUpper,
	}

	// Create and parse a template named "funcExample".
	// The template uses the custom function "toUpper" to convert input text to uppercase.
	tmpl := template.Must(template.New("funcExample").Funcs(funcs).Parse("Hello, {{. | toUpper}}!"))
	//.Funcs(funcs): Registers the function map so that "toUpper" can be used inside the template.
	//.Parse("Hello, {{. | toUpper}}!"): Defines the template content, where {{. | toUpper}} applies toUpper to the input data

	// Execute the template with "golang" as input data.
	// The input string is passed through the "toUpper" function.
	tmpl.Execute(os.Stdout, "golang")

	// Expected output:
	// Hello, GOLANG!
}
