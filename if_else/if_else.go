package main

import "fmt"

func main() {
	var age int = 108
	if age == 5 {
		fmt.Println(true)
	} else if age > 23 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
	// variable can declear inside if construct
	if age := 10; age == 10 {
		fmt.Println("age is 10")
	}
}
