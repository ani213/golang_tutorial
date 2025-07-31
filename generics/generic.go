package main

import "fmt"

func printSlice[T any](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

func printSlice2[T interface{}](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

func printSlice3[T int | string](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}
func printSlice4[T comparable](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

func printSlice5[T comparable, V string](arr []T, name V) {
	for _, item := range arr {
		fmt.Println(item, name)
	}
}

type Stack[T any] struct {
	elements []T
}

func main() {
	intValue := []int{1, 2, 3, 4, 5}
	stringValue := []string{"Hello", "World", "Generics"}
	floatValue := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	printSlice(intValue)
	printSlice(stringValue)
	printSlice(floatValue)
	printSlice([]struct{ name string }{{name: "Alice"}, {name: "Bob"}, {name: "Charlie"}})
	printSlice([]map[string]int{{"one": 1}, {"two": 2}, {"three": 3}})
	printSlice2([]int{1, 2, 3, 4, 5})
	printSlice2([]string{"Generics", "are", "powerful"})
	printSlice2([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	printSlice2([]bool{true, false, true})
	printSlice2([]struct{ name string }{{name: "Alice"}, {name: "Bob"}, {name: "Charlie"}})
	printSlice2([]map[string]int{{"one": 1}, {"two": 2}, {"three": 3}})
	fmt.Println("Generics in Go are powerful and flexible, allowing you to write reusable code that works with any data type.")
	printSlice3(intValue)
	printSlice4(intValue)
	myStack := Stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}
	fmt.Println("Stack elements:", myStack.elements)
	printSlice5([]int{1, 2, 3}, "Generics in Go")
}
