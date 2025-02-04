package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler function to retrieve and display parameters from the route
func NameHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// params will be [{firstname John} {lastname Doe}]
	fmt.Fprintf(w, "Name: %s %s", p.ByName("firstname"), p.ByName("lastname"))
}

func main() {
	// Create a new router instance
	router := httprouter.New()

	// Define a route with dynamic parameters
	// Params are not optional here
	// NameHandler will be called if the url path matches, ex- name/John/Doe
	router.GET("/name/:firstname/:lastname", NameHandler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", router)
}
