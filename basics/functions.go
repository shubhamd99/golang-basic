package main

import "fmt"

func main() {
	msg := "Hello User"
	writeMessage(msg)
	fmt.Println(msg)

	// Passing the address
	msg2 := "Shubham"
	writeMessage2(&msg)
	fmt.Println(msg2)

	// Multiple args
	total := sum("Sum: ", 1, 2, 3, 4)
	fmt.Println(total)

	// Error handling in functions
	value, err := divide(3, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	// Anonymous function
	func() {
		fmt.Println("This is Anonymous function")
	}()

	// Assign Anonymous function
	// We cannot call this function before its declaration
	fun := func() {
		fmt.Println("This is Anonymous function")
	}
	fun()

	// Calling Methods
	rect := rectangle{
		width:  20,
		height: 10,
	}

	area := rect.area()
	fmt.Println("Area of rectangle: , ", area)
}

// Rectangle Struct
type rectangle struct {
	width, height int
}

// Declaration of the Methods for Rectangle struct
func (r rectangle) area() int {
	return r.width * r.height
}

func writeMessage(msg string) {
	msg = "Empty"
	fmt.Println(msg)
}

func writeMessage2(msg *string) {
	*msg = "Hello Linkedin" // DeReference before assigning
	fmt.Println(*msg)
}

func sum(msg string, values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
