package main

import (
	"database/sql" // Import the database/sql package for working with databases
	"fmt"          // Import fmt for printing output

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver to establish a connection with MySQL
)

// Person struct to map database records into Go objects
type Person struct {
	FirstName string `json:"first_name"` // JSON field mapping
	LastName  string `json:"last_name"`  // JSON field mapping
}

func main() {
	// Step 1: Connect to MySQL database
	// The connection string format: "username:password@protocol(address)/dbname"
	db, err := sql.Open("mysql", "root:password@123@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return // Exit the program if the connection fails
	}
	defer db.Close() // Ensure the database connection is closed when function exits

	// Step 2: Create a table if it does not already exist
	_, err = db.Query("CREATE TABLE IF NOT EXISTS PERSON (first_name VARCHAR(50), last_name VARCHAR(50))")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return // Exit the program if table creation fails
	}

	// Step 3: Prepare an SQL statement for inserting data into the table
	stmt, err := db.Prepare("INSERT INTO PERSON(first_name, last_name) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return // Exit if preparing the statement fails
	}
	defer stmt.Close() // Ensure the prepared statement is closed after execution

	// Step 4: Execute the prepared statement with sample data
	_, err = stmt.Exec("Bob", "Barker") // Insert first person
	if err != nil {
		fmt.Println("Error inserting first record:", err)
		return
	}

	_, err = stmt.Exec("Betty", "White") // Insert second person
	if err != nil {
		fmt.Println("Error inserting second record:", err)
		return
	}

	// Step 5: Retrieve data from the PERSON table
	dataset, err := db.Query("SELECT * FROM PERSON") // Query all records
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer dataset.Close() // Ensure dataset is closed after processing

	// Step 6: Loop through the query results and process each row
	// In Go, range is typically used for iterating over slices, arrays, maps, and channels,
	// but it cannot be used directly with *sql.Rows (which is returned by db.Query()).
	for dataset.Next() {
		var person Person
		// Scan the retrieved row data into the `Person` struct fields
		err := dataset.Scan(&person.FirstName, &person.LastName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue // Skip to the next record if scanning fails
		}
		// Print the retrieved person record in Go struct format
		fmt.Println(person)
	}

	// Step 7: Handle any errors that may have occurred during iteration
	if err = dataset.Err(); err != nil {
		fmt.Println("Error processing dataset:", err)
	}
}
