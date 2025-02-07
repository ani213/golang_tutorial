package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver for database connection
	"github.com/jmoiron/sqlx"          // Import sqlx package for easier SQL handling
)

// Person struct represents a row in the PERSON table
type Person struct {
	FirstName string `db:"first_name"` // Maps to the "first_name" column in the database
	LastName  string `db:"last_name"`  // Maps to the "last_name" column in the database
}

func main() {
	// Establish a connection to the MySQL database using sqlx
	db, err := sqlx.Connect("mysql", "root:password@123@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close() // Ensure the database connection is closed when the function exits

	// SQL query to create the PERSON table if it does not already exist
	query := `
	CREATE TABLE IF NOT EXISTS PERSON (
		first_name VARCHAR(50),  -- Column for first name
		last_name VARCHAR(50)    -- Column for last name
	)`
	_, err = db.Exec(query) // Execute the SQL query
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// SQL query to insert records using named parameters for better readability
	insertQuery := `INSERT INTO PERSON (first_name, last_name) VALUES (:first_name, :last_name)`

	// List of people to insert into the database
	people := []Person{
		{"Peter", "Parker"},
		{"Laira", "Doe"},
	}

	// Loop through each person and insert into the database using NamedExec
	for _, person := range people {
		_, err := db.NamedExec(insertQuery, person) // Automatically maps struct fields to SQL parameters
		if err != nil {
			log.Println("Error inserting data:", err)
		}
	}

	// Declare a slice to store the retrieved records
	var persons []Person

	// Fetch all records from the PERSON table into the persons slice
	err = db.Select(&persons, "SELECT * FROM PERSON")
	if err != nil {
		log.Fatal("Error fetching records:", err)
	}
	// fmt.Println(persons)
	// output -
	// [{Peter Parker} {Laira Doe}]

	// Iterate over the retrieved records and print them
	for _, person := range persons {
		fmt.Println(person) // Prints each person's details
		// 	output -
		// 	 {Peter Parker}
		//   {Laira Doe}
	}
}

// sqlx-
// -With sqlx, you can fetch all records into a slice in one line
// -With sqlx, it automatically maps NULL values to string fields without issues

// Standard database/sql-
// -Fetching multiple rows is manual using Standard database/sql (Looping over Rows)
// -Standard database/sql requires sql.NullString for handling NULL values:

// Summary: Why Use sqlx? 🚀

//     Feature	           database/sql	               sqlx
// Struct Mapping	❌ Manual Scan() required	✅ StructScan()
// Fetching Rows	❌ Loop + Scan()         	✅ Select(&slice, query)
// Named Queries	❌ Positional args (?)	    ✅ NamedExec() with struct fields
// Single Record	❌ QueryRow().Scan()	        ✅ Get(&struct, query)
// NULL Handling	❌ Needs sql.NullString   	✅ Auto-converts NULLs
// Less Boilerplate	❌ More manual work	        ✅ Cleaner & simpler code
