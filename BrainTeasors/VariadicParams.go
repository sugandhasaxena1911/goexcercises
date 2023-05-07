package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4} //sum i 10
	n := VariadicParamV1(sl...)
	fmt.Println("The sum is ", n)
	fmt.Println("the slice after the function call is ", sl) // Remains same

	//V2
	sl2 := []int{1, 2, 3, 4, 5} //sum i 15

	n = VariadicParamV2(sl2...)
	fmt.Println("The sum is ", n)
	fmt.Println("the slice after the function call is ", sl2)

}

func VariadicParamV1(a ...int) int {
	sum := 0
	a = append(a, []int{1, 2, 3}...)
	a = append(a, 1, 2, 3)
	for _, v := range a {
		sum = sum + v
	}

	return sum

}
func VariadicParamV2(a ...int) int {
	sum := 0

	a[1] = 100
	for _, v := range a {
		sum = sum + v
	}

	return sum

}
