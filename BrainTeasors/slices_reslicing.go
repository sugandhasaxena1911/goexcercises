package main

import "fmt"

func main() {

	var a []int
	fmt.Println(a, len(a), cap(a))
	// POINT 1 : How capacity varies of the underlying array
	fmt.Println("---------------------------")
	for x := 1; x <= 10; x++ {
		a = append(a, x)
		fmt.Println(a, len(a), cap(a))

	}

	fmt.Println("playing with Re-slicing, underlying array is same")
	fmt.Println("---------------------------")
	fmt.Println(a, len(a), cap(a)) //10,16
	a = a[5:8]
	fmt.Println(a, len(a), cap(a)) //3,11
	a = a[4:9]
	fmt.Println(a, len(a), cap(a)) //5,7

	a = a[:4]
	fmt.Println(a, len(a), cap(a)) //4,7

	a = a[:7]
	fmt.Println(a, len(a), cap(a)) //7,7

}
