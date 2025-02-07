package main

import "fmt"

func main() {
	mapValue := make(map[string]int)
	mapValue["x"] = 2
	mapValue["y"] = 3
	fmt.Println("mapvale", mapValue)
	fmt.Println("x:", mapValue["x"])
}
