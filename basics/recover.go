package main

import "fmt"

// This function is created to handle
// the panic occurs in entry function
// but it does not handle the panic
// occurred in the entry function
// because it called in the normal
// function
func handlepanic() {

	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

// Function
func entry(lang *string, aname *string) {

	// Normal function
	handlepanic()

	// When the value of lang
	// is nil it will panic
	if lang == nil {
		panic("Error: Language cannot be nil")
	}

	// When the value of aname
	// is nil it will panic
	if aname == nil {
		panic("Error: Author name cannot be nil")
	}

	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
	fmt.Printf("Return successfully from the entry function")
}

// Main function
func main() {

	A_lang := "GO Language"
	entry(&A_lang, nil)
	fmt.Printf("Return successfully from the main function")
}
