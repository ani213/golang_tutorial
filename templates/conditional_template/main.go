package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Admin bool
}

func main() {
	tmpl := template.Must(template.New("checkAdmin").Parse("User: {{.Name}} {{if .Admin}}(Admin){{else}}(Regular User){{end}}\n"))

	users := []User{
		{Name: "Alice", Admin: true},
		{Name: "Bob", Admin: false},
	}

	for _, user := range users {
		tmpl.Execute(os.Stdout, user)
	}
	// output -
	// User: Alice (Admin)
	// User: Bob (Regular User)

}
