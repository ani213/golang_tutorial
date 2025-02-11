package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Person struct represents a person with an ID, First Name, and Last Name
type Person struct {
	ID        int    `json:"id"`    // ID field, mapped to "id" in JSON
	FirstName string `json:"fname"` // FirstName field, mapped to "fname" in JSON
	LastName  string `json:"lname"` // LastName field, mapped to "lname" in JSON
}

func main() {

	// Slice to store the list of people (acts as an in-memory database)
	var people []Person

	// Handle HTTP requests at "/person/" endpoint
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		// Handle GET request - Fetch all persons
		case http.MethodGet:
			// Convert the slice of people into a JSON-formatted response
			jsonResponse, _ := json.MarshalIndent(&people, " ", "   ")
			fmt.Fprint(w, string(jsonResponse)) // Send the response

		// Handle POST request - Add a new person
		case http.MethodPost:
			// Read the request body
			requestBodyBytes, _ := io.ReadAll(r.Body)

			// Unmarshal (convert) JSON data into a new Person struct
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			// Check if the ID already exists
			for i := range people {
				if people[i].ID == newPerson.ID {
					w.WriteHeader(http.StatusConflict) // Return HTTP 409 Conflict if ID exists
					return
				}
			}

			// Append the new person to the slice
			people = append(people, newPerson)

		// Handle PUT request - Update an existing person
		case http.MethodPut:
			// Read the request body
			requestBodyBytes, _ := io.ReadAll(r.Body)

			// Convert JSON into a Person struct
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			// Find the person with the given ID and update their details
			for i := range people {
				if people[i].ID == newPerson.ID {
					people[i].FirstName = newPerson.FirstName
					people[i].LastName = newPerson.LastName
					return
				}
			}

			// If ID not found, return HTTP 404 Not Found
			w.WriteHeader(http.StatusNotFound)

		// Handle DELETE request - Remove a person from the list
		case http.MethodDelete:
			// Read the request body
			requestBodyBytes, _ := io.ReadAll(r.Body)

			// Convert JSON into a Person struct
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			// Find the person with the given ID and delete them
			for i := range people {
				if people[i].ID == newPerson.ID {
					copy(people[i:], people[i+1:])  // Shift elements left to remove the person
					people = people[:len(people)-1] // Resize the slice
					return
				}
			}

			// If ID not found, return HTTP 404 Not Found
			w.WriteHeader(http.StatusNotFound)

		// Handle any unsupported HTTP methods
		default:
			w.WriteHeader(http.StatusMethodNotAllowed) // Return HTTP 405 Method Not Allowed
			return
		}
	})

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
