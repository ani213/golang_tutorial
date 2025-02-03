package main

import (
	"net/http" //It imports the net/http package to handle HTTP requests.
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/unrecognized_path.html")
	})
	// The http.HandleFunc function is used to define different routes and serve corresponding HTML files.
	// If a user accesses /, the server responds with the file unrecognized_path.html.
	// This acts as a catch-all for unrecognized paths.

	http.HandleFunc("/demopath1/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1.html")
	})

	http.HandleFunc("/demopath1/subpatha", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1_subpatha.html")
	})

	http.HandleFunc("/demopath2/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath2.html")
	})

	http.ListenAndServe(":8080", nil)
	// The nil parameter means it will use the default HTTP request multiplexer (http.DefaultServeMux).
}
