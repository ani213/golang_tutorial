package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("list").Parse("List of Names:\n{{range .}}{{.}}\n{{end}}"))
	// A static string "List of Names:\n".
	// A range loop that iterates over the input data (.) and prints each item ({{.}}), followed by a newline.
	// The loop ends with {{end}}.

	names := []string{"Alice", "Bob", "Charlie"}

	tmpl.Execute(os.Stdout, names)
	// output -
	// List of Names:
	// Alice
	// Bob
	// Charlie

}
