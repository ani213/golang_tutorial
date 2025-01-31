package main

import "fmt"

func main() {
	safeFunction()
	fmt.Println("Program continues after panic recovery.")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil { // recover return value which is passed in panic.
			fmt.Println("Recovered from panic:", r)
		}
	}()
	y := 4
	fmt.Println("Before panic")
	// panic("Something went wrong!") // This panic will be recovered
	x := [3]int{1, 2, 3}
	x[y] = 5
	fmt.Println("After panic") // This line will not execute
}
