package main

import "fmt"

/*
A variadic function is a function that accepts a variable number of arguments.
 If the last parameter of a function definition is prefixed by ellipsis …,
 then the function can accept any number of arguments for that parameter.
*/
func main() {
	// variadicFun(1, 2, 3, 4, 5, 6)
	variadicFun2(7, 8, 9, 10, 11, 12, 13)
}

// func variadicFun(a int, b ...int) {
// 	x := b
// 	fmt.Println("x", x)
// 	fmt.Println("a", a)
// 	for i, v := range x {
// 		fmt.Println("index:", i, "  Value:", v)
// 	}

// }

func variadicFun2(a int, b ...int) {
	if len(b) < 4 {
		fmt.Println("Not enough values in b to unpack")
		return
	}

	x, y, z, d := b[0], b[1], b[2], b[3:]
	fmt.Println("x:", x)
	fmt.Println("a:", a)

	fmt.Println("Values in d slice:")
	for i, v := range d {
		fmt.Println("index:", i, "  Value:", v)
	}

	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d) // d is a slice of remaining elements
}
