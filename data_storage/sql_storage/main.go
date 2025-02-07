package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "root:password@123#edc@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Create table if it doesn't exist
	_, err = db.Query("CREATE TABLE IF NOT EXISTS PERSON (first_name VARCHAR(50), last_name VARCHAR(50))")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Insert data using prepared statements
	stmt, err := db.Prepare("INSERT INTO PERSON(first_name, last_name) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	stmt.Exec("Bob", "Barker")
	stmt.Exec("Betty", "White")

	// Query data from the table
	dataset, err := db.Query("SELECT * FROM PERSON")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer dataset.Close()

	// Loop through results and print them
	for dataset.Next() {
		var person Person
		err := dataset.Scan(&person.FirstName, &person.LastName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Println(person)
	}
}
