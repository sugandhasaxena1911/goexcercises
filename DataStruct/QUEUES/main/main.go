package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/DataStruct/QUEUES"
)

func main() {
	Queue1 := QUEUES.Queue{}
	for x := 1; x <= 5; x++ {
		Queue1.Enqueue(x)
	}
	Queue1.PrintQueue()
	fmt.Println("SIZE", Queue1.Size())
	fmt.Println(Queue1.IsEmpty())
	fmt.Println(Queue1.Dequeue())
	fmt.Println("AFTER DEQUEUE")
	Queue1.PrintQueue()

}
