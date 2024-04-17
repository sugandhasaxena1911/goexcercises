package main

import "fmt"

func main() {
	stack1 := NewStack()
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	fmt.Println(stack1.Empty())
	fmt.Println(stack1.Top())
	stack1.Traverse()
	fmt.Println(stack1.Pop())
	fmt.Println(stack1.Pop())
	stack1.Traverse()

}

type Stack struct {
	top int
	st  []int
}

func NewStack() *Stack {
	return &Stack{-1, []int{}}

}

func (s *Stack) Push(val int) {
	s.st = append(s.st, val)
	s.top++

}
func (s *Stack) Pop() int {
	if s.top == -1 {
		fmt.Println("Stack underflow ")
		return s.top
	}
	fmt.Println("Popped ")

	val := s.st[s.top]
	s.top--
	return val

}

func (s *Stack) Top() int {
	if s.top == -1 {
		fmt.Println("Stack underflow ")
		return s.top
	}
	return s.st[s.top]

}

func (s *Stack) Empty() bool {
	if s.top == -1 {
		return true
	}

	return false

}

func (s *Stack) Traverse() {

	for x := s.top; x >= 0; x-- {
		fmt.Println(s.st[x])
	}
}
