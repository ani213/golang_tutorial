package main

import "fmt"

// name:= "Aniket"  can not use out side of function
const name = "Last name"

func main() {
	const name1 string = "Aniket"
	fmt.Println(name)

	const name2 = "Anurag"

	fmt.Println(name1)
	fmt.Println(name2)
	const (
		firstName = "Aniket"
		lastName  = "Kumar"
		port      = 8080
		host      = "localhost:"
	)
	fmt.Println(firstName, lastName, port, host)
}
