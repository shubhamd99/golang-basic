package main

import (
	"fmt"
	"strconv"
)

var i = 32 // Global Variable (Camel casing, not exported)
var I = 32 // If we create variable starting with capital then it is exported to other package (Pascal casing)

func main() {
	// We cannot declare variable with same name in same scope
	var i int // integer variable
	i = 12 // assign integer

	var j int = 14 // we can do "var j = 14" also here, go will automatically detect the data type

	var m = float32(i)

	var d float64 // Initialized by zero by default

	h := 12 // short form

	var sToInt string = strconv.Itoa(h) // Int to String

	var boolVal bool = true

	fmt.Println(sToInt, boolVal) // 12

	fmt.Println(i, j, h, m, d)
	fmt.Printf("%v, %T\n", i, i) // 12 int (value and type)


	// Arithmetic
	sum1 := 20 + 3 // * - / %
	fmt.Println(sum1)

	// Binary Operators
	bin1 := 2 & 2
	// bin1 := 2 & 2 AND
	// bin1 := 2 | 2 OR
	// bin1 := 2 ^ 2 caret
	// bin1 := 2 &^ 2 caret &
	fmt.Println(bin1)

	var comp complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", comp, comp) // (1+2i), complex128

	str1 := "Shubham"
	str2 := "Dhage"

	fmt.Println(str1 + " " + str2)
	fmt.Printf("%v, %T\n", string(str1[2]), str1) // u, string


	// Convert string into byte array
	byStr := []byte(str2) // [68 104 97 103 101]
	fmt.Println(byStr)

	// Constants
	const con int = 12
	fmt.Println(con)

	// Group constants
	const (
		User string = "Shubham"
		Product string = "Samsung"
	)

	// The iota keyword represents successive integer constants 0, 1, 2,…
	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)
	const (
		C3 = iota // it will work in constant block itself
		C4
		C5
	)
	fmt.Println(C0, C1, C2) // "0 1 2"
	fmt.Println(C3, C4, C5) // "0 1 2"
}