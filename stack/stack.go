package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {

	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {

	l := len(s.items) - 1
	removedItems := s.items[l]
	s.items = s.items[:l]
	return removedItems
}

func main() {

	myStack := Stack{}
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
