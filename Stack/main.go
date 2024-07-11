package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}


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

	
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Push(4)

	fmt.Print(stack.items,"\n\n")

	for {

     item,err:= stack.Pop()

	  if err!=nil{
		fmt.Println(err)
		break
	  }else{
		fmt.Println(item)
	  }

	//   fmt.Println()

	}

}
