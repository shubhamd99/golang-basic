package main

import (
	"fmt"
	"time"
)

func portal1(channel1 chan string) {
	time.Sleep(1 * time.Millisecond)
	channel1 <- "Welcome to Channel 1"
}

func portal2(channel2 chan string) {
	time.Sleep(2 * time.Millisecond)
	channel2 <- "Welcome to Channel 2"
}

func main() {

	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	// Select is like Switch statement for channels
	// Only the first channel will run that will have less wait time
	select {
	case opt1 := <-R1:
		fmt.Println(opt1)
	case opt2 := <-R2:
		fmt.Println(opt2)
	}
}
