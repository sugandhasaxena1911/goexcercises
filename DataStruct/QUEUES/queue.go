package QUEUES

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(value int) {
	node := Node{value, nil}
	if q.IsEmpty() {
		q.Tail = &node
		q.Head = &node
	} else {
		q.Tail.Next = &node
		q.Tail = &node
	}

}
func (q *Queue) Dequeue() int {
	n := 0
	if q.IsEmpty() {
		return -1
	} else if q.Size() == 1 {
		n = q.Head.Value
		q.Head = nil
		q.Tail = nil
		return n

	} else {
		n = q.Head.Value
		q.Head = q.Head.Next
		return n

	}

}

func (q *Queue) IsEmpty() bool {
	if q.Head == nil {
		return true
	} else {
		return false
	}

}
func (q *Queue) Size() int {
	n := 0
	if q.IsEmpty() {
		return n
	} else {
		currentnode := q.Head
		for currentnode != nil {
			n++
			currentnode = currentnode.Next
		}
		return n
	}
}

func (q *Queue) Peek() *Node {
	if q.IsEmpty() {
		return nil
	} else {
		return q.Head

	}

}
func (q *Queue) PrintQueue() {

	currentnode := q.Head
	for currentnode != nil {
		fmt.Println(currentnode.Value)
		currentnode = currentnode.Next
	}

}
