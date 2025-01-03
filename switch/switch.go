package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println(i)
	default:
		fmt.Println("Default")
	}
	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it's workday")
	}
	// interface{ } - means i can anytype
	whoAmI := func(i interface{}) {
		switch i.(type) { // i.(type) it's return type of i
		case int:
			fmt.Println("it is an Integer")
		case string:
			fmt.Println("it is string")
		case bool:
			fmt.Println("it is boolean")
		default:
			fmt.Println("its other")
		}
	}
	whoAmI(1)
	whoAmI("golang")
	whoAmI(true)
	whoAmI(5.89)

}
