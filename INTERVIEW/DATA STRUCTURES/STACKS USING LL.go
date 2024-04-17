package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}
type Stacks struct {
	top *Node
}

func (s *Stacks) Push(val int) {
	node := Node{val, nil}
	if s.IsEmpty() {
		s.top = &node
		return
	}

	current := s.top
	node.next = current
	s.top = &node

}
func (s *Stacks) Pop() int {
	if s.IsEmpty() {
		fmt.Println("cant pop em , Empty Stack")
		return -1
	}
	val := s.top.val
	s.top = s.top.next

	return val

}
func (s *Stacks) TraverseStacks() {
	fmt.Println("printing stacks ")
	current := s.top
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}
func (s *Stacks) IsEmpty() bool {

	if s.top == nil {
		return true
	} else {
		return false
	}
}

func main() {
	s := Stacks{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.TraverseStacks()
	s.Pop()
	s.TraverseStacks()

}
