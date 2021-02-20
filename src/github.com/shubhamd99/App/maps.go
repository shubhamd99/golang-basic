package main

import "fmt"

func main() {

	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"Laptop":   1000,
	}
	fmt.Println(shoppingCart, len(shoppingCart))

	shoppingCart["Keyboard"] = 1200
	fmt.Println(shoppingCart["Keyboard"])

	// With make method
	shoppingCart2 := make(map[string]int)
	shoppingCart2 = map[string]int{
		"Shubham": 100,
		"Rohan":   50,
		"Meesho":  1000,
	}
	fmt.Println(shoppingCart2)

	// Availability of the key
	// cart, ok := shoppingCart["Mobile"] // 0, false
	_, ok := shoppingCart["Mobile"]
	fmt.Println(ok) // false

	// Copy data, it will copy memory location ,so changing any one will change the parent map also
	sc := shoppingCart
	fmt.Println("sc:", sc)

	// Delete Key in Map
	delete(shoppingCart, "Keyboard")
	fmt.Println(shoppingCart)
}
