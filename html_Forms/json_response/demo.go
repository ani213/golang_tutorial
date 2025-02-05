package main

import (
	"encoding/json"
	"net/http"
)

// Person struct with default JSON marshalling
type Person struct {
	FirstName string
	LastName  string
}

// PersonJsonFields struct with custom JSON field names
type PersonJsonFields struct {
	FirstName string `json:"fname"` // JSON output will have "fname" instead of "FirstName"
	LastName  string `json:"lname"` // JSON output will have "lname" instead of "LastName"
}

// PersonJsonMarshaler struct with a custom JSON marshaler
type PersonJsonMarshaler struct {
	FirstName string
	LastName  string
}

// Custom JSON marshaling for PersonJsonMarshaler
func (p *PersonJsonMarshaler) MarshalJSON() ([]byte, error) {
	// Manually format the JSON output as {"Name":"FirstName LastName"}
	return []byte("{\"Name\":\"" + p.FirstName + " " + p.LastName + "\"}"), nil
}

func main() {

	// Default JSON serialization using struct field names
	http.HandleFunc("/json/default/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&Person{FirstName: "Bob", LastName: "Barker"})
		w.Header().Set("Content-Type", "application/json") // Set JSON content type
		w.Write(j)
	})
	// output -
	// {
	//   "firstName": "Bob",
	//   "lastName": "Barker"
	// }

	// Custom JSON field names using struct tags
	http.HandleFunc("/json/fields/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&PersonJsonFields{FirstName: "Bob", LastName: "Barker"})
		w.Header().Set("Content-Type", "application/json") // Set JSON content type
		w.Write(j)
	})
	// output -
	// {
	//   "fname": "Bob",
	//   "lname": "Barker"
	// }

	// Custom JSON marshaling with a combined "Name" field
	http.HandleFunc("/json/marshaler/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&PersonJsonMarshaler{FirstName: "Bob", LastName: "Barker"})
		w.Header().Set("Content-Type", "application/json") // Set JSON content type
		w.Write(j)
	})
	// The MarshalJSON method overrides the default behavior,
	// combining first and last name into a single "Name" field.
	// output -
	// {
	//   "Name": "Bob Barker"
	// }

	// Bypasses Struct Tags → This method ignores struct tags like json:"fname" because it manually constructs the JSON.
	// This technique is useful when you need a custom JSON format different from the standard Go struct-to-JSON conversion.

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
