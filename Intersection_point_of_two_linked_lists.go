package main

import "fmt"

// create singly linked list
type Node struct {
	value any
	next  *Node
}

type SLL struct {
	head Node
}

type SLLInterface interface {
	display()
}

func add(sll *SLL, value any) {

	// create head
	if sll.head.next == nil {
		sll.head.value = value
		sll.head.next = nil
		return
	}

	currentNode := sll.head
	var newNode = Node{value: value, next: nil}
	for currentNode.next != nil {
		currentNode = *currentNode.next
		if currentNode.next != nil {
			currentNode.next = &newNode
			currentNode = newNode
		}
	}
}

func NewSLL(values ...any) SLL {
	var sll SLL
	for _, value := range values {
		fmt.Printf("Appending %v->", value)
		add(&sll, value)
		sll.display()
	}
	return sll
}

func (sll *SLL) display() {
	fmt.Println("\nSLL : ", sll)
	currentNode := sll.head
	for currentNode.next != nil {
		fmt.Printf("%v->", currentNode.value)
		currentNode = *currentNode.next
	}
}

func main() {
	// create 2 sll
	// 1. sll
	sll1 := NewSLL(1, 2, 3, 6, 5, 6)
	sll1.display()
	// sll2 := NewSLL(30, 31, 32, 35, 37, 43)
	// sll2.display()

	// make intersection of them
}
