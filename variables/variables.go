package main

import "fmt"

func main() {
	var name string = "Golang"
	fmt.Println(name)

	// go lang auto infer
	var name1 = "Golang"
	fmt.Println(name1)

	// int
	var age int = 5
	fmt.Println(age)

	// bool
	var isAdult bool = true
	fmt.Println(isAdult)

	var price = 4.7
	fmt.Println((price))

	// shorthand syntex
	nameShortHand := "Golang"
	fmt.Println((nameShortHand))
	// assign after
	var lastName string
	lastName = "Verma"
	fmt.Println(lastName)
}
