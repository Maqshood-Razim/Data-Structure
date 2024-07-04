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

func (list *LinkedList) add(values ...int) {

	for _, values := range values {
		newNode := &Node{value: values}

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

func(list *LinkedList) display(){

    current := list.head

	for current!=nil{
		fmt.Print(current.value,"->")
		current=current.next
	}

	fmt.Println("nil")


}

func main() {
	list := LinkedList{}
	list.add(3, 5, 19)
	list.display()
}