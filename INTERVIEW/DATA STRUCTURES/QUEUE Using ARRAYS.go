package main

import "fmt"

type Queue struct {
	front int
	rear  int
	arr   []int
}

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	q.TraverseQueue()
	q.Dequeue()
	q.TraverseQueue()

	q.Dequeue()
	q.Dequeue()
	q.TraverseQueue()
	q.Dequeue()
	q.TraverseQueue()

	q.Dequeue()
	q.Enqueue(1)
	q.TraverseQueue()

}

func NewQueue() *Queue {
	return &Queue{-1, -1, []int{}}

}

func (q *Queue) Enqueue(val int) {
	if q.front == -1 && q.rear == -1 {
		q.front = 0
		q.rear = 0
	} else {
		q.rear++
	}
	q.arr = append(q.arr, val)

}
func (q *Queue) Dequeue() int {
	// No element
	if q.front == -1 && q.rear == -1 {
		fmt.Println("Empty queue, cant dequeue")
		return -1
	}
	// if only one element
	if q.front == q.rear {
		val := q.arr[q.front]
		q.front = -1
		q.rear = -1
		return val
	}
	val := q.arr[q.front]
	q.front++
	return val
}
func (q *Queue) TraverseQueue() {
	if q.front == -1 && q.rear == -1 {
		fmt.Println("Empty queue")
		return
	}
	fmt.Println("PRINTING QueueS \n")
	for x := q.front; x <= q.rear; x++ {
		fmt.Println(q.arr[x])
	}
}
