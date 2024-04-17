package main

import "fmt"

func main() {
	Queue1 := NewQueue()
	Queue1.Enqueue(1)
	Queue1.Enqueue(2)
	Queue1.Enqueue(3)
	fmt.Println(Queue1.Front())
	Queue1.Traverse()
	Queue1.Dequeue()

	Queue1.Traverse()

}

type Queue struct {
	front int
	rear  int
	q     []int
}

func NewQueue() *Queue {
	return &Queue{0, 0, []int{}}

}

func (qu *Queue) Enqueue(val int) {
	qu.q = append(qu.q, val)
	qu.rear++

}
func (qu *Queue) Dequeue() int {
	if qu.rear == 0 {
		fmt.Println("empty quuee")
		return -1
	}
	val := qu.q[qu.front]
	qu.shiftqueue()

	return val

}
func (qu *Queue) shiftqueue() {
	if qu.rear != 1 {

		for x := 0; x >= qu.rear-1; x++ {
			qu.q[x] = qu.q[x+1]

		}
	}
	qu.rear--

}

func (qu *Queue) Front() int {
	if qu.rear == 0 {
		fmt.Println("Empty queue")
		return -1
	}
	val := qu.q[qu.front]
	return val
}

func (qu *Queue) Traverse() {
	if qu.rear == 0 {
		fmt.Println("Empty queue")
		return
	}
	fmt.Println("The QUEUE 
	")
	for x := qu.front; x < qu.rear; x++ {
		fmt.Println(qu.q[x])
	}
}
