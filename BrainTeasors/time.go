package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 3000
	fmt.Println("world")
	//	time.Sleep(timeout * time.Millisecond)    // int and duration mismatch

	time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("hello")
}
