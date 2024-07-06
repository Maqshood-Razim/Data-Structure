package main

import (
	"fmt"
)

// Stack struct represents a stack that holds a slice of integers
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

// Pop removes and returns the top item of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	topIndex := len(s.items) - 1
	topItem := s.items[topIndex]
	s.items = s.items[:topIndex]
	return topItem, nil
}

func main() {
	stack := Stack{}

	// Push elements onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Push(4)

	fmt.Print(stack.items,"\n\n")

	for i := 0; i < 2; i++ {
		item, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(item)
		}
	}

}
