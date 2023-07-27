package STACKS

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	top *Node
}

func (st *Stack) PrintStack() {
	currentnode := st.top
	for currentnode != nil {
		fmt.Println(currentnode.Value)
		currentnode = currentnode.Next

	}
}

func (st *Stack) Push(node *Node) {
	currentnode := node
	if st.IsEmpty() {
		st.top = currentnode
	} else {
		previousnode := st.top
		currentnode.Next = previousnode
		st.top = currentnode

	}
}

func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		previous := st.top.Next
		st.top = previous

		return true

	}

}

func (st *Stack) Peek() *Node {

	return st.top

}

func (st *Stack) Size() int {
	n := 0
	if st.IsEmpty() {
		return n
	}
	currentnode := st.top
	for currentnode != nil {
		n++
		currentnode = currentnode.Next
	}
	return n
}
func (st *Stack) IsEmpty() bool {
	if st.top == nil {
		return true
	} else {
		return false
	}
}
