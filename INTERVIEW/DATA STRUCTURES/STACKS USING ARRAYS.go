package main

import "fmt"

type Stack struct {
	top int
	arr []int
}

func main() {
	st := NewStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	st.TraverseStack()
	fmt.Println(st.Pop())
	st.TraverseStack()

}

func NewStack() *Stack {
	return &Stack{-1, []int{}}

}

func (st *Stack) Push(val int) {
	st.arr = append(st.arr, val)
	st.top++
}
func (st *Stack) Pop() int {
	if st.top == -1 {
		fmt.Println("Stack is empty , cant Pop")
		return -1
	}
	val := st.arr[st.top]
	st.top--
	return val
}
func (st *Stack) TraverseStack() {
	if st.top == -1 {
		fmt.Println("Stack is empty ")
		return

	}
	fmt.Println("PRINTING STACKS \n")
	for x := st.top; x >= 0; x-- {
		fmt.Println(st.arr[x])
	}
}
