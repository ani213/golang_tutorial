package main

import (
	"fmt"
	"os"
)

func closer(f *os.File) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	f.Close()
	fmt.Println(f.Name(), "successfully closed")
	return nil
}

func main() {
	// f, err := os.Create("create.txt")
	file, _ := os.Open("create.txt")

	closer(file)

	fmt.Println(file, "hyuow")
}
