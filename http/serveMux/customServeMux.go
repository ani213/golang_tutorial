package main

import (
	"fmt"
	"log"
	"net/http"
)

// Custom handler function
func CustomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a custom handler.")
}

// Middleware to log requests
func LogRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("About to serve request.")
		// here handler == serveMux
		handler.ServeHTTP(w, r)
		log.Println("Request served.")
	})
}

func main() {
	// Create a new custom ServeMux (multiplexer)
	serveMux := http.NewServeMux()

	// Register the handler for "/custom/"
	serveMux.HandleFunc("/custom/", CustomHandler)

	// Here we can use http as default serve mux here but default serve mux is global and accessible outside the function.
	// http.HandleFunc("/custom/", CustomHandler)
	// but due to security reason it is good to use custom serve mux

	// Wrap the ServeMux with the logging middleware
	loggedServeMux := LogRequests(serveMux)

	// Start the HTTP server on port 8080
	log.Println("Server is running on port 8080...")
	// These lines start the HTTP server on port 8080, using the loggedServeMux as the handler.
	http.ListenAndServe(":8080", loggedServeMux)
}
