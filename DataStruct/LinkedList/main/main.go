package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/DataStruct/LinkedList"
)

func main() {
	
	var list1 LinkedList.LinkedList // list1 is nil
	list1.AddNode(5)
	fmt.Println(list1)
	fmt.Println(list1.Head.Value, list1.Head.Next)
	list1.AddNode(17)
	list1.AddNode(18)
	list1.AddNode(2)
	list1.AddNode(9)
	list1.AddNode(6)
	list1.TraverseList()
	list1.DeleteList(5)
	list1.TraverseList()
	list1.DeleteList(6)
	list1.TraverseList()
	list1.DeleteList(2)
	list1.TraverseList()
	list1.DeleteList(18)
	list1.TraverseList()

}
