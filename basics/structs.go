package main

import "fmt"

// Student struct
type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {

	student1 := Student{
		name:   "Shubham Dhage",
		rollNo: 5,
		subjects: []string{
			"Maths",
			"Physics",
			"chemistry",
		},
	}

	var student2 Student = Student{
		"Rohan Singh",
		10,
		[]string{
			"Maths",
			"Physics",
		},
	}

	var student3 Student = Student{
		name: "John Smith",
	}

	// Structs are value type, it will only pass the information about the value, it wont pass the memory pointer location
	student4 := student1
	// student4 := &student1 // --> To send the pointer information instead of value
	student4.name = "Ayush gupta"

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)
	fmt.Println(student4)

	fmt.Println(student1.name, student1.subjects[2])

	student1.rollNo = 10
	fmt.Println(student1)
}
