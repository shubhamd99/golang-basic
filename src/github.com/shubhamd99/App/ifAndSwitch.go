package main

import "fmt"

func main() {

	if true {
		fmt.Println("This is simple if statement")
	}

	// Created i variable and assigned 2
	if i := 2; i == 3 {
		fmt.Println("This is simple if statement block")
	} else if i == 2 {
		fmt.Println("This is simple else if statement block")
	} else {
		fmt.Println("This is simple else statement block")
	}

	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"Shubham": 100,
		"Rohan":   50,
		"Swiggy":  1000,
	}

	// Assigning the Swiggy key present bool in ok and checking whether the key is present or not in the map
	if _, ok := shoppingCart["Swiggy"]; ok {
		fmt.Println("Item exist in the shopping cart")
	}

	i := 10
	j := 20

	if i > 0 && j > 0 {
		fmt.Println("i, j is greater than 0")
	}

	if i >= 0 || j >= 0 {
		fmt.Println("i, j is greater than 0")
	}

	// Switch statement
	switch 3 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")

	}

	// Expression and statement
	switch o := 2 + 3; o {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 8:
		fmt.Println("Even")
	default:
		fmt.Println("default")

	}

	switch {
	case j > 0:
		fmt.Println("Big")
	case j < 0:
		fmt.Println("Small")
	default:
		fmt.Println("default")

	}

	// Type checking with switch
	var k interface{} = 5.3
	switch k.(type) {
	case int:
		fmt.Println("This is int")
		break // Go will take care of break but if we want we can add
		fmt.Println("This is int")
	case float64:
		fmt.Println("This is float64")
	case string:
		fmt.Println("This is string")
	default:
		fmt.Println("This is default")
	}

	// Fallthrough
	// Output: 1 , 2
	switch 1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Default")

	}
}
