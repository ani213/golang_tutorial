package main

import (
	"fmt"
)

// Person struct represents an individual with a first and last name
type Person struct {
	FirstName string
	LastName  string
}

// Group struct represents a team that contains multiple people
type Group struct {
	Name   string
	People []Person
}

func main() {
	// Creating Person instances
	person1 := Person{"George", "Smith"}
	person2 := Person{"Elizabeth", "Platch"}
	person3 := Person{"Ronald", "Duthers"}
	person4 := Person{"Lucy", "Brown"}

	// Creating slices of people
	people1 := []Person{person1, person2} // First group
	people2 := []Person{person3, person4} // Second group

	// Creating Group instances and storing them in a slice
	gp := []Group{
		{Name: "Accounting", People: people1},
		{Name: "Sales", People: people2},
	}

	// Printing the groups and their details
	fmt.Println(gp)                       // Display the entire array of groups
	fmt.Println(gp[0])                    // Display the first group (Accounting)
	fmt.Println(gp[1])                    // Display the second group (Sales)
	fmt.Println(gp[1].People[0])          // Display the first person in the second (Sales) group
	fmt.Println(gp[1].People[0].LastName) // Display the last name of the first person in the Sales group
}

// output -
// [{Accounting [{George Smith} {Elizabeth Platch}]} {Sales [{Ronald Duthers} {Lucy Brown}]}]
// {Accounting [{George Smith} {Elizabeth Platch}]}
// {Sales [{Ronald Duthers} {Lucy Brown}]}
// {Ronald Duthers}
// Duthers
