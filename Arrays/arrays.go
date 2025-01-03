package main

import "fmt"

func main() {
	// zeroed values
	var nums [4]int // default value added 0 for int and flot
	nums[0] = 9
	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)
	var str [4]string // empty string assing for string i/o []
	fmt.Println(str)
	var boolean [4]bool // default value added false
	fmt.Println(boolean)
	// to declear in a single line
	num := [3]int{1, 2, 3}
	fmt.Println(num)

	// 2d array

	array2d := [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(array2d)

	// - Fixed size, that is predictable
	// - Memory optimazation
	// - constant time access
}
