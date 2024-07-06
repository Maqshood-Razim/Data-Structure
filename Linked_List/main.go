package main

import (
	"fmt"
)

func main() {

	list := Linkedlist{}

	fmt.Println("enter the values")

	var val int

	for i := 0; i < 7; i++ {
		fmt.Scan(&val)
		list.Add(val)
	}

	fmt.Print("linked list : ")
	list.Display()

}

type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	head *Node
}

func (list *Linkedlist) Add(val ...int) {

	for _, val := range val {
		newNode := &Node{value: val}

		if list.head == nil {
			list.head = newNode
		} else {

			current := list.head

			for current.next != nil {
				current = current.next
			}
			current.next = newNode

		}
	}

}

func (list *Linkedlist) Display() {

	current := list.head

	for current != nil {

		fmt.Print(current.value, "->")
		current = current.next

	}

	fmt.Println("nil")

}
