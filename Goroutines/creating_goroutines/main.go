package main

import (
	"fmt"
	"time"
)

// myFunc prints numbers from 1 to 100 with a given label
func myFunc(step string) {
	for i := 1; i < 50; i++ {
		fmt.Print(step, ": ", i, " \n") // Prints step name along with the current number
	}
}

func main() {
	// Call myFunc synchronously
	myFunc("synchronous1")
	// Start myFunc as a goroutine (runs concurrently)
	go myFunc("G1")
	go myFunc("G2")
	myFunc("synchronous2")
	go myFunc("G3")
	go myFunc("G4")
	myFunc("synchronous3")
	// Allow some time for goroutines to execute before the program exits
	time.Sleep(100 * time.Millisecond) //not a good method to pause, use waitgroup instead

	fmt.Println("finished") // Print when main function is about to exit
}
