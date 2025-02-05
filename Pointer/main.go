package main

import "fmt"

// pointer default value is <nil>
func main() {
	// var ptr *int
	// fmt.Println("var value pointer", ptr) // o/p <nil>

	myNumber := 1
	var ptr = &myNumber

	fmt.Println("vlue of ptr", ptr)
	fmt.Println("vlue of ptr", *ptr)

	*ptr = *ptr + 1
	fmt.Println("mynumber value", myNumber)

	modifiedMynumberValueByAddress(&myNumber)
	fmt.Println("value is", myNumber)
}

func modifiedMynumberValueByAddress(num *int) {
	*num = *num + 1
}
