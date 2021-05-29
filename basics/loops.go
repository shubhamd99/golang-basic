package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	k := 0
	for ; k < 10; k++ {
		fmt.Println(k)
	}

	t := 0
	for {
		fmt.Println("t", t)

		t = t + 1

		if t == 6 {
			break
		}
	}

	f := 0
	for f < 10 {
		f = f + 1
		if f == 6 {
			continue // skip
		}

		fmt.Println("f", f)
	}

	// Nested Loops
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " * ", j, "=", i*j)
		}
	}

	// Loop over Arrays
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	for key, value := range arr {
		fmt.Println(key, value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}

	// Loop over Maps
	shoppingCart2 := make(map[string]int)
	shoppingCart2 = map[string]int{
		"Shubham": 100,
		"Rohan":   50,
		"Meesho":  1000,
	}

	for k, v := range shoppingCart2 {
		fmt.Println(k, v)
	}

	for k := range shoppingCart2 {
		fmt.Println(k)
	}
}
