package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get command-line arguments
	args := os.Args

	// Ensure at least two arguments are provided (port and message)
	if len(args) < 3 {
		fmt.Println("Usage: go run main.go <port> <message>")
		os.Exit(1) // Exit with an error code
	}

	port := args[1]    // First argument is the port number
	message := args[2] // Second argument is the response message

	// Print the arguments for debugging purposes
	fmt.Println("Arguments:", args)

	// Define an HTTP handler for the root ("/") route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message)) // Write the message as a response
	})

	// Start the HTTP server on the specified port
	fmt.Println("Starting server on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

//how to run - go run main.go 8080 "Hello, World!"
