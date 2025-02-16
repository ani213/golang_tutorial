package main

import (
	"encoding/xml"
	"fmt"
)

// Person struct represents a person with XML annotations
type Person struct {
	FirstName string   `xml:"fname"`   // Maps to <fname> in XML
	LastName  string   `xml:"lname"`   // Maps to <lname> in XML
	Hobbies   []string `xml:"hobbies"` // Maps to <hobbies> in XML
}

func main() {
	// Create an instance of Person and initialize values
	joe := Person{FirstName: "Joe", LastName: "Smith"}
	joe.Hobbies = []string{"Skiing", "Wind Surfing"}

	// Marshal the struct into an XML formatted string with indentation
	xmlOutput, _ := xml.MarshalIndent(joe, " ", "  ")

	// Print the XML header and formatted XML output
	fmt.Println(xml.Header) // Prints XML header <?xml version="1.0" encoding="UTF-8"?>
	// output -
	// 	<?xml version="1.0" encoding="UTF-8"?>
	fmt.Println(string(xmlOutput)) // Prints the formatted XML representation of the struct
	// output -
	//  <Person>
	//    <fname>Joe</fname>
	//    <lname>Smith</lname>
	//    <hobbies>Skiing</hobbies>
	//    <hobbies>Wind Surfing</hobbies>
	//  </Person>

	// Declare an empty Person struct to store unmarshalled data
	joeFromXml := Person{}

	// Unmarshal XML back into the struct
	xml.Unmarshal(xmlOutput, &joeFromXml)

	// Print the unmarshalled struct
	fmt.Println(joeFromXml)
	// output -
	// {Joe Smith [Skiing Wind Surfing]}

}
