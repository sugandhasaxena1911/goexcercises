package LinkedList

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	Head *Node
}

func (list *LinkedList) AddNode(n int) {
	//empty list
	NewNode := &Node{n, nil}
	if list.Head == nil {
		list.Head = NewNode
		return
	}
	//list is not empty
	// need to reach the end of list and add Value
	fmt.Println("Reached here ")
	currentNode := list.Head
	fmt.Println("Reached here ", currentNode.Value, currentNode.Next)

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	//here current node is last node
	currentNode.Next = NewNode
	fmt.Println("Reached here 3", currentNode.Value, currentNode.Next)
	fmt.Println("Reached here 4", currentNode.Value, currentNode.Next)

}

func (list *LinkedList) TraverseList() {
	fmt.Println("The elements of list are ")
	currentNode := list.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func (list *LinkedList) DeleteList(n int) {
	if list.Head == nil {
		return
	}
	//first element match
	if list.Head.Value == n {
		list.Head = list.Head.Next
		return
	}

	currentNode := list.Head.Next
	previousNode := list.Head
	for currentNode != nil {
		if currentNode.Value == n {
			previousNode.Next = currentNode.Next
			return
		}
		previousNode = currentNode
		currentNode = currentNode.Next

	}
}
