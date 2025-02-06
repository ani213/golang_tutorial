package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int = 567
	// int to float
	var x float64 = float64(number)
	fmt.Println("x", x)
	// int to String
	var str string = strconv.Itoa(number)
	fmt.Println("String", str)
	// String to int
	intvalue, _ := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("error", err)
	// }
	fmt.Println("int", intvalue)
	// String to float
	var floatString string = "3.14"
	floatValue, _ := strconv.ParseFloat(floatString, 10)
	fmt.Println("floatValue", floatValue)

}
