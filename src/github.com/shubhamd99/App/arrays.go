package main

import "fmt"

func main() {

	var amounts [3]int = [3]int{10, 20, 30}
	fmt.Printf("%v, %T\n", amounts, amounts) // [10 20 30], [3]int

	amt := [3]int{40, 50, 60}  // short form
	fmt.Println(amt, len(amt)) // [40 50 60] 3

	arrayWithUnknownSize := [...]int{30, 40, 50, 60, 70}
	fmt.Println(arrayWithUnknownSize, len(arrayWithUnknownSize), arrayWithUnknownSize[3])

	// Change array elements
	amounts[0] = 20

	// Copy array
	copyArr := amounts
	copyArr[0] = 30
	fmt.Println(copyArr, amounts) // [30 20 30] [20 20 30]

	// Copy memory location instead of array content
	copyMemArr := &amounts
	copyMemArr[0] = 100
	fmt.Println(copyMemArr, amounts) // &[100 20 30] [100 20 30]

	// Copy whole content
	a := [...]int{20, 30, 40, 50, 70, 80, 100, 30, 101}
	b := a[:]
	fmt.Printf("A: %v\n", a) // A: [20 30 40 50]
	fmt.Printf("B: %v\n", b) // B: [20 30 40 50]

	// Slicing operations
	c := a[2:]              // 2nd index to last index
	d := a[:5]              // 0th index to 5th index
	e := a[2:7]             // 2nd index to 7th index
	f := a[:len(a)-1]       // Remove last element
	fmt.Println(c, d, e, f) // [40 50 70 80 100 30 100] , [20 30 40 50 70] , [40 50 70 80 100]

	// Matrix array
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}

	fmt.Println(identityMatrix)       // [[1 0 0] [0 1 0] [0 0 1]]
	fmt.Println(identityMatrix[1][1]) // 1

	identityMatrix[1][1] = 10
	fmt.Println(identityMatrix[1][1]) // 10
}
