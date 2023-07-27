package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/DataStruct/STACKS"
)

func main() {
	stack1 := STACKS.Stack{}
	for x := 0; x <= 5; x++ {
		node1 := STACKS.Node{Value: x}
		stack1.Push(&node1)
	}
	stack1.PrintStack()
	stack2 := STACKS.Stack{}
	fmt.Println("Is empty ?", stack2.IsEmpty())
	fmt.Println(stack1.Size())
	fmt.Println(stack2.Size())

	fmt.Println("POP", stack1.Pop())
	fmt.Println("After POP")
	fmt.Println(stack1.Size())
	stack1.PrintStack()

	node2 := stack1.Peek()
	fmt.Println("The peeked element ", node2.Value)

}
