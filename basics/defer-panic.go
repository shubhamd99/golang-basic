package main

import (
	"fmt"
	"os"
)

func main() {
	// Defer
	// 1
	// 3
	// 2
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	// Panic
	// fmt.Println(1)
	// panic("This is panic")
	// fmt.Println(3)

	// Create File functionality
	f := createFile("defer.txt")
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		// Panic is a built-in function that stops the ordinary flow of control and begins panicking.
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data is there")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
