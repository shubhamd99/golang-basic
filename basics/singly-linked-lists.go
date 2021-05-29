package main

import "fmt"

// ------------ Structs -----------

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node // DeReference
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node // DeReference
	len  int
}

// ------------ Functions -----------

func main() {

	l := LinkedList{}

	fmt.Println("\n************* Insert *************")
	l.Insert(2)
	l.Insert(4)
	l.Insert(3)

	fmt.Println("************* Print *************")
	l.Print()
}

// ------------ Interfaces -----------

// Insert inserts new node at the end of  from linked list
func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n // Allocating memory pointer
		l.len++
		return
	}

	ptr := l.head
	for i := 0; i < l.len; i++ {
		// If next pointer is null then will insert new list item
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}

	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}
