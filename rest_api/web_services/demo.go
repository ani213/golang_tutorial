package main

import (
	"fmt"      // For formatted I/O
	"io"       // For reading from the request body
	"net/http" // For creating the HTTP server and handling requests
)

func main() {
	// Handle the root URL "/" with a simple handler that returns a 400 Bad Request.
	// This could be used as a default or catch-all endpoint.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write an HTTP status code 400 (Bad Request) to the response header.
		w.WriteHeader(http.StatusBadRequest)
		// No body is written in this handler.
	})

	// Handle the "/withargs/" URL.
	// This handler reads query parameters "arg1" and "arg2" from the URL.
	http.HandleFunc("/withargs/", func(w http.ResponseWriter, r *http.Request) {
		// r.Method is the HTTP method (GET, POST, etc.)
		// r.URL.Query() returns a map of query parameters.
		// We use Get("arg1") and Get("arg2") to retrieve the values for these parameters.
		response := fmt.Sprintf("This is a response from a %s request. The arguments are: arg1:%s, arg2:%s",
			r.Method,
			r.URL.Query().Get("arg1"),
			r.URL.Query().Get("arg2"))

		// Write the formatted response string to the ResponseWriter.
		fmt.Fprint(w, response)
	})

	// Handle the "/postonlywithbody/" URL.
	// This handler is designed to process POST requests only and will echo back the body of the request.
	http.HandleFunc("/postonlywithbody/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the HTTP method is POST.
		if r.Method != http.MethodPost {
			// If it's not a POST, respond with 405 Method Not Allowed.
			w.WriteHeader(http.StatusMethodNotAllowed)
			return // Exit the handler if the method is not POST.
		}

		// Read the entire request body.
		// Note: The error returned by io.ReadAll is ignored here; in production, proper error handling is advised.
		body, _ := io.ReadAll(r.Body)

		// Convert the body bytes to a string and write it to the response.
		// This effectively echoes back the request body.
		fmt.Fprintln(w, string(body))
	})

	// Start the HTTP server on port 8080.
	// The second parameter is nil, meaning we are using the default HTTP ServeMux.
	http.ListenAndServe(":8080", nil)
}
