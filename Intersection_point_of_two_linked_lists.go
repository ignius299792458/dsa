package main

import "fmt"

// create singly linked list
type Node struct {
	value any
	next  *Node
}

type SLL struct {
	head *Node
	size int
}

type SLLInterface interface {
	display()
	add(value any)
}

func (sll *SLL) add(value any) {
	newNode := &Node{value: value, next: nil}

	if sll.head == nil {
		sll.head = newNode
		sll.size++
		return
	}

	currentNode := sll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	} // till currentNode.next is nil

	currentNode.next = newNode
	sll.size++
}

func NewSLL(values ...any) *SLL {
	sll := &SLL{head: nil, size: 0}
	for _, value := range values {
		sll.add(value)
	}
	return sll
}

func (sll *SLL) display() {
	for currentNode := sll.head; currentNode.next != nil; {
		fmt.Printf(" %v|%v ->", currentNode.value, currentNode.next)
		currentNode = currentNode.next
	}
	fmt.Println("\nSize: ", sll.size)
}

func main() {
	sll1 := NewSLL(1, 2, 3, 4, 5, 6)
	sll1.display()
	sll2 := NewSLL(7, 8, 9, 10, 11)
	sll2.display()
}
