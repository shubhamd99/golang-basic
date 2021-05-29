package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg) // pointer of the wait group
	}

	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // It will execute at the end of the function

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
