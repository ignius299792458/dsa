package main

// create singly linked list
type Node struct {
	value any
	next  *Node
}

var head Node

type SinglyLLInterface interface {
	AddToLink(value any)
}

func (n *Node) AddToLink(value any) {

	// create head
	if head.value == nil && head.next == nil {
		head.value = value
		head.next = nil
		return
	}

	prevNode := head
	var newNode Node
	prevNode.next = &newNode
	for prevNode.next != nil {

	}

}

// create 2 sll

// make intersection of them
