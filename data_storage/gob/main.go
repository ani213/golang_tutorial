package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Person struct defines a simple structure with a first and last name.
type Person struct {
	FirstName string
	LastName  string
}

// writeBinaryFile serializes the given data using GOB encoding and writes it to a file.
func writeBinaryFile(data interface{}, file string) error {
	// Create a buffer to hold the encoded data
	buf := new(bytes.Buffer)

	// Create a new GOB encoder
	encoder := gob.NewEncoder(buf)

	// Encode the data into the buffer
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encoding failed: %v", err)
	}

	// Write the encoded data to a file using os.WriteFile (Go 1.16+)
	err := os.WriteFile(file, buf.Bytes(), 0600)
	if err != nil {
		return fmt.Errorf("file write failed: %v", err)
	}
	return nil
}

// readBinaryFile reads a binary GOB file and decodes it into the provided data structure.
func readBinaryFile(data interface{}, file string) error {
	// Read the binary file content using os.ReadFile (Go 1.16+)
	raw, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("file read failed: %v", err)
	}

	// Create a buffer from the read data
	buf := bytes.NewBuffer(raw)

	// Create a new GOB decoder
	decoder := gob.NewDecoder(buf)

	// Decode the binary data into the provided data structure
	if err := decoder.Decode(data); err != nil {
		return fmt.Errorf("decoding failed: %v", err)
	}
	return nil
}

func main() {
	// Define the file path for storing the binary data
	file := "person.gob"

	// Create a new Person instance
	person := Person{"Bob", "Barker"}

	// Write the Person instance to a binary file
	if err := writeBinaryFile(person, file); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	// Create an empty Person struct to hold the decoded data
	var newPerson Person

	// Read and decode the binary file back into the Person struct
	if err := readBinaryFile(&newPerson, file); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print the decoded struct to verify successful serialization and deserialization
	fmt.Println("Decoded Struct:", newPerson)
}
