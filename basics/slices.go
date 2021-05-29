package main

import "fmt"

func main() {
	// Slices have internal array application, we can increase the size of the array
	var a []int = []int{1, 2, 3} // slice now
	var b []int = a              // it will point out the pointer array location
	b[0] = 4

	fmt.Println(a) // [4 2 3]
	fmt.Println(b) // [4 2 3]

	fmt.Println(len(a), cap(a)) // length - 3, capacity - 3

	var d []int = append(a, 5)
	fmt.Println(d) // [4 2 3 5]

	var e []int = append(a[1:], 5) // remove first elm and push 5
	fmt.Println(e)                 // [2 3 5]

	// data of a is appended to e and e is assigned to f
	var f []int = append(e, a...)
	fmt.Println(f) // [2 3 5 4 2 3]

	// Create slice with main method
	var c []int = make([]int, 3, 10) // size - 3, capacity - 10
	fmt.Println(len(c), cap(c))

}
