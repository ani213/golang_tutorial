package main

import (
	"encoding/json" // Importing encoding/json for JSON encoding and decoding
	"fmt"           // Importing fmt for printing output
)

// Person struct represents a person with first name, last name, and hobbies
type Person struct {
	FirstName string   `json:"fname"`   // JSON key for FirstName is "fname"
	LastName  string   `json:"lname"`   // JSON key for LastName is "lname"
	Hobbies   []string `json:"hobbies"` // JSON key for Hobbies is "hobbies"
}

func main() {
	// Creating a new instance of Person
	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"Skiing", "Wind Surfing"} // Assigning hobbies

	// MarshalIndent is used to convert struct to formatted JSON with indentation
	// &joe: Pointer to joe struct
	// " ": Prefix (empty in this case)
	// "  ": Indentation (2 spaces per level)
	jsonOutput, _ := json.MarshalIndent(&joe, " ", "  ")

	// Print the JSON output as a string
	fmt.Println(string(jsonOutput))
	// output -
	// {
	// 	"fname": "Joe",
	// 	"lname": "Smith",
	// 	"hobbies": [
	// 	  "Skiing",
	// 	  "Wind Surfing"
	// 	]
	//   }

	// Creating an empty struct to hold unmarshalled JSON data
	joeFromJson := Person{}

	// Unmarshal JSON data back into the struct
	json.Unmarshal(jsonOutput, &joeFromJson)

	// Print the struct to verify JSON decoding
	fmt.Println(joeFromJson)
	// output -
	// {Joe Smith [Skiing Wind Surfing]}

}
