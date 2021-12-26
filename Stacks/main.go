package main

import "fmt"

// hold slice
type Stack struct {
	items []int
}

// push
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// pop
func (s *Stack) Pop() int {
	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length]
	return toRemove
}

func main() {
	myStack := Stack{}

	fmt.Println(myStack)
	myStack.Push(40)
	myStack.Push(50)
	myStack.Push(60)
	myStack.Push(70)
	myStack.Push(80)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

}
