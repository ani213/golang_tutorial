package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("main").Parse(`
{{define "header"}}--- Header Section ---{{end}}
{{define "footer"}}--- Footer Section ---{{end}}

{{template "header"}}
Hello, World!
{{template "footer"}}
`))
	// {{define "header"}}--- Header Section ---{{end}}: Defines a section called "header" with the content --- Header Section ---.
	// {{define "footer"}}--- Footer Section ---{{end}}: Defines a section called "footer" with the content --- Footer Section ---.
	// {{template "header"}}: This renders the content of the "header" section.
	// Hello, World!: This is the central content of the template.
	// {{template "footer"}}: This renders the content of the "footer" section.

	tmpl.Execute(os.Stdout, nil)
	// output -
	// --- Header Section ---
	// Hello, World!
	// --- Footer Section ---

}
