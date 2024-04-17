package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(val int) {
	node := Node{val, nil}
	if q.IsEmpty() {
		q.front = &node
		q.rear = &node
		return
	}
	current := q.rear
	current.next = &node
	q.rear = &node

}
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("cant dequeue , Empty Queue")
		return -1
	}
	val := q.front.val
	current := q.front
	q.front = current.next
	return val

}
func (q *Queue) TraverseQueue() {
	fmt.Println("printing Queue ")
	current := q.front
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}
func (q *Queue) IsEmpty() bool {

	if q.front == nil && q.rear == nil {
		return true
	}
	return false
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.TraverseQueue()

	q.Dequeue()
	q.TraverseQueue()

}
