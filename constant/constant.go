package main

import "fmt"

// name:= "Aniket"  can not use out side of function
const name = "Last name"

func main() {
	const name string = "Aniket"
	fmt.Println(name)

	const name2 = "Aniket"

	fmt.Println(name2)
	fmt.Println(name2)
	const (
		firstName = "Aniket"
		lastName  = "Kumar"
		port      = 8080
		host      = "localhost:"
	)
	fmt.Println(firstName, lastName, port, host)
}
