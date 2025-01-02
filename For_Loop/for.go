package main

import "fmt"

func main() {
	// while implement
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i, "new")
	}
	for v := range 10 {
		fmt.Println(v)
	}
}
