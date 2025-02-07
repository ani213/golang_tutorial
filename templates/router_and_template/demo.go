package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {

	tmpl := "<B><font color='green'>First Name: </font></B> {{.FirstName}} <BR> <B><font color='red'>Last Name: </font></B> {{.LastName}}"

	router := httprouter.New()

	router.GET("/name/:firstName/:lastName", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		t := template.New("NameTemplate")
		tp, _ := t.Parse(tmpl)

		tp.Execute(w, &Person{ps.ByName("firstName"), ps.ByName("lastName")})
	})

	http.ListenAndServe(":8080", router)
}
