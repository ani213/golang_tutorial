package main

import "fmt"

// slices is dynamic array
// most used construct
// usefull methods
func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len((nums)))

	var slice = make([]int, 2, 5) // 2 is initial lenngth and 5 is capacity
	/* capacity-> maxium number of element can be fit ( it's not mean
	only 5 element can fit, this is dynamic array so capacity will be increase automatic
	if element will add more) */
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)

	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	fmt.Println(slice)

}
