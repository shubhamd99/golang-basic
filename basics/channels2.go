package main

import (
	"fmt"
	"sync"
)

var waitGrp1 = sync.WaitGroup{}

// Buffer Channel
func main() {
	// This channel will always send and recieve data in int format
	ch := make(chan int)

	waitGrp1.Add(2) // Two Goroutine func

	// Read Channel
	// Restrict this to only Receive the data
	go func(ch <-chan int) {
		// Receiving the data from channel

		/*for rangeOfCh := range ch {
			fmt.Println(rangeOfCh)
		}*/

		for {
			// If data is avaialble in the channel we will process it
			if rangeOfCh, ok := <-ch; ok {
				fmt.Println(rangeOfCh)
			} else {
				break
			}
		}

		waitGrp1.Done()
	}(ch)

	// Write Channel
	// Restrict this to only Pass the data
	go func(ch chan<- int) {
		// Passing the data to the channel
		ch <- 24
		ch <- 100
		ch <- 30
		close(ch) // We have to close the channel

		waitGrp1.Done()
	}(ch)

	waitGrp1.Wait()
}
