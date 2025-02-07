package main

import "fmt"

func main() {
	// 1st using make function
	mapValue := make(map[string]int)
	mapValue["x"] = 2
	mapValue["y"] = 3
	fmt.Println("mapvale", mapValue)
	fmt.Println("x:", mapValue["x"])

	// 2nd
	var newMap = map[string]string{
		"1": "2",
	}
	newMap["2"] = "3"
	fmt.Println("newMap", newMap)

	// Checking if a key exists

	value, ok := mapValue["x"]
	fmt.Println("value", value, "isExist", ok)

	value1, ok := mapValue["not"]
	fmt.Println("value", value1, "isExist", ok)

	value2, ok := newMap["not"]
	fmt.Println("value", value2, "isExist", ok)
}
