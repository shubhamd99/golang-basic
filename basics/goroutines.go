package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// Concurrent application
func main() {

	// A goroutine is a lightweight thread of execution.
	// go writeMessage1() // It will finish with main method

	message := "Shubham"

	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // Done the waiting of wait group
	}(message)

	// Dirty delay for main method
	// time.Sleep(100 * time.Millisecond)

	// We have to tell the go compiler that we are using a goroutines over here and you need to wait until goroutine completed
	wg.Wait()
}

func writeMessage1() {
	fmt.Println("Hello")
}
