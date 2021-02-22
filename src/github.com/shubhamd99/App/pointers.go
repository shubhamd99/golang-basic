package main

import "fmt"

type Foo struct {
	bar int
}

func main() {

	var a int = 12
	var b *int = &a // memory of a

	fmt.Println(a) // 12
	fmt.Println(b) // 0xc0000b2008
	// DeReference
	fmt.Println(*b) // 12

	// Pointer with struct
	var foo *Foo
	fmt.Println(foo) // <nil>
	foo = new(Foo)
	fmt.Println(foo) // &{0}
	// DeReference
	(*foo).bar = 300
	// foo.bar = 300 // we can write this also golang will take care of DeReference internally
	fmt.Println((*foo).bar) // 0
}
