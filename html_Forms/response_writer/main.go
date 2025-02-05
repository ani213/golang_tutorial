package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/plaintext/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "<h1>This is plain text.</h1>")
	})

	http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "<h1>This is HTML.</h1>")
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static.html")
	})

	http.ListenAndServe(":8080", nil)
}
