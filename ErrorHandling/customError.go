package main

import (
	"errors"
	"fmt"
	"math"
)

func volume(r float64) (float64, error) {
	if r < 0 {
		return 0, errors.New("redious is negative value")
	}
	return (4 / 3) * math.Pi * r * r * r, nil
}
func main() {
	v, err := volume(-4.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%0.4f", v) // 201.0619>= (%0.4f) 4 decimal digit
}
