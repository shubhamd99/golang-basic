package main

import "fmt"

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memoryType     string
}

// Embedded relationship (composition relationship)
type Computer struct {
	Processor
	Memory
	price int
}

func main() {

	computer := Computer{}
	computer.price = 50000
	computer.memoryCapacity = 3000
	computer.cores = 4

	computer1 := Computer{
		Processor: Processor{
			processorName: "Intel i7",
			cores:         8,
		},
		Memory: Memory{
			memoryCapacity: 8,
			memoryType:     "DDR4",
		},
		price: 70000,
	}

	fmt.Println(computer)  // {{ 4} {3000 } 50000}
	fmt.Println(computer1) // {{Intel i7 8} {8 DDR4} 70000}

}
