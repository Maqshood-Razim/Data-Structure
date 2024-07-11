package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{value: value, next: l.head}
	l.head = newNode
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}

}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}
func main() {
	l := &LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Print("After appending: ")
	l.Display()

	l.Prepend(0)
	l.Prepend(19)
	fmt.Print("After prepending: ")
	l.Display()

	l.DeleteWithValue(2)
	fmt.Print("After deleting\n")
	l.Display()

}
