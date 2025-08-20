package main

/*
------------------------------------------------
********* Example : 1 *********
---
   A:  1 → 2 → 3
                 ↘
                   7 → 8 → null
                 ↗
   B:       4 → 5

------------------------------------------------
********* Example : 2 *********
---
List A:  1 → 2 → 3 → 4 → null
List B:  1 → 2 → 3 → 4 → null
------------------------------------------------
*/

import (
	"fmt"
)

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
	intersectionPointWith(sll *SLL) *Node
	interMixPointRef(sll *SLL)
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

// intersectionPointWith
func (thisSLL *SLL) intersectionPointWith(sll *SLL) *Node {

	outerCurrentNode := thisSLL.head
	for outerCurrentNode.next != nil {

		// check if same reference found in given sll
		for innerCurrentNode := sll.head; innerCurrentNode.next != nil; {
			if outerCurrentNode.next == innerCurrentNode.next {
				return outerCurrentNode.next // this is the intersection point
			}
			innerCurrentNode = innerCurrentNode.next
		}
		outerCurrentNode = outerCurrentNode.next
	}
	return nil
}

func main() {
	sll1 := NewSLL(1, 2, 3, 6, 7)
	sll1.display()
	sll2 := NewSLL(4, 5, 6, 7)
	sll2.display()

	// intersection point of sll1 and sll2
	fmt.Println("\n->>> intersection point of sll1 and sll2: ", sll1.intersectionPointWith(sll2))
}
