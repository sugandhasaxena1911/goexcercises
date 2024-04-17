package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(val int) {
	node := Node{val, nil}
	if ll.Head == nil {
		ll.Head = &node
		return
	}

	current := ll.Head
	for current.next != nil {
		current = current.next
	}
	current.next = &node

}

func (ll *LinkedList) DeleteFromEnd() int {
	if ll.Head == nil {
		fmt.Println("Empty List , cannot delete ")
		return -1
	}
	if ll.Head.next == nil {
		val := ll.Head.val
		ll.Head = nil
		return val
	}

	current := ll.Head
	var previous *Node
	for current.next != nil {
		previous = current
		current = current.next

	}
	val := current.val
	previous.next = nil
	return val

}
func (ll *LinkedList) DeleteByval(val int) {
	if ll.Head == nil {
		fmt.Println("Empty List , cannot delete ")
		return
	}
	if ll.Head.next == nil && ll.Head.val == val {
		ll.Head = nil
		return
	}

	current := ll.Head
	var previous *Node
	for current != nil {
		if current.val == val {
			if current == ll.Head {
				ll.Head = ll.Head.next
				return
			}
			previous.next = current.next
			return
		}
		previous = current
		current = current.next

	}
}
func (ll *LinkedList) TraverseLInkedList() {
	fmt.Println("Printing Linked List")

	current := ll.Head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}

func (ll *LinkedList) ReverseLL() {
	current := ll.Head
	var previous *Node
	var forward *Node

	for current != nil {
		forward = current.next
		current.next = previous
		previous = current
		current = forward

	}
	ll.Head = previous
}
func main() {

	ll := LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.TraverseLInkedList()

	ll.DeleteByval(2)
	ll.TraverseLInkedList()

	ll.DeleteFromEnd()
	ll.TraverseLInkedList()

	ll.Add(5)
	ll.Add(6)
	ll.TraverseLInkedList()
	ll.ReverseLL()
	ll.TraverseLInkedList()

}
