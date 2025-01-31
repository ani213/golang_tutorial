package main

import "fmt"

func main() {
	fmt.Println("hi")
	defer fmt.Println("world")

	fmt.Println(add())
}

func add() int {
	defer fmt.Println("1")
	defer fmt.Println("2")
	return 5
}
