package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int
	mu   sync.Mutex // Mutex to protect the view count
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Unlock the mutex before exiting
		wg.Done()     // Signal that this goroutine is done
	}()
	p.mu.Lock()
	p.view++
	// p.mu.Unlock()                               // Unlock the mutex after incrementing
	fmt.Println("Incremented view to:", p.view) // Print the current view count
	// Simulate some work
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		view: 0,
	}
	for range 200 {
		wg.Add(1)
		go myPost.increment(&wg) // Increment the view count concurrently
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(myPost.view)
}
