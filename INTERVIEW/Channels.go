package main

import "fmt"

func main() {

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go sendOddEven(odd, even, quit)

	// case 1 :using range
	/*
		go func() {
			for x := range odd {
				fmt.Println("Odd ", x)
			}

		}()

		for y := range even {
			fmt.Println("Even ", y)
		}
		fmt.Println("All odd even printed ")
	*/
	/*  case 2 using for select */

	for {
		select {
		case v := <-odd:
			fmt.Println("odds", v)
		case v := <-even:
			fmt.Println("evens", v)
		case v := <-quit:
			fmt.Println("All odds evens printed ", v)
			return
		}

	}
}

func sendOddEven(odd, even chan int, quit chan int) {

	for x := 1; x <= 100; x++ {

		if x%2 == 0 {
			even <- x
		} else {
			odd <- x
		}

	}

	//close(odd)   case1
	//close(even)  case 1

	quit <- 1
}
