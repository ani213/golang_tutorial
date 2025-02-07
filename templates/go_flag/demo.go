package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// Define command-line flags for the port and response message
	var port, message string

	flag.StringVar(&port, "port", "8080", "The HTTP port to listen on") // Default port: 8080
	flag.StringVar(&message, "message", "", "The response message")     // No default message
	flag.Parse()                                                        // Parse the flags

	// Inform the user about the server details
	fmt.Printf("Starting server on port %s...\n", port)
	fmt.Printf("Response message: %q\n", message)

	// Define an HTTP handler for the root ("/") route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message)) // Send the response message
	})

	// Start the HTTP server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
