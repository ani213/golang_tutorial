package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func sum1(a, b int) int {
	return a + b
}

func moreThen2return(a, c, b int) (int, int, int) {
	return a, b, c
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	fn := func(a int) int {
		return a
	}
	processIt(fn)
	fn1 := processIt2()
	fmt.Println(fn1(2))
	fmt.Println(sum(3, 5))
	fmt.Println(sum1(3, 5))
	a, b, _ := moreThen2return(3, 5, 6)
	fmt.Println(a, b)

}
