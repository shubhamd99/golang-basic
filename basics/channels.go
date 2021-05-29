package main

import (
	"fmt"
	"sync"
)

var waitGrp = sync.WaitGroup{}

func main() {
	// This channel will always send and recieve data in int format
	ch := make(chan int)

	waitGrp.Add(2) // Two Goroutine func
	go func() {
		// Receiving the data from channel
		i := <-ch
		fmt.Println(i) // 12

		// Passing the data to the channel
		ch <- 27

		waitGrp.Done()
	}()
	go func() {
		// Passing the data to the channel
		temp := 12
		ch <- temp
		temp = 24

		// Receiving the data from channel
		fmt.Println(<-ch)

		waitGrp.Done()
	}()

	waitGrp.Wait()
}
