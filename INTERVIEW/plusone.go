package main

import "fmt"

func main() {
	digits := []int{9, 9}
	fmt.Println(plusOne(digits))

}

// 1299 1300   189 190   199 200   194  195,  99 100
func plusOne(digits []int) []int {
	//chk := false
	n := len(digits)
	for x := n - 1; x >= 0; x-- {
		if digits[x] != 9 {
			digits[x] = digits[x] + 1
			return digits
		} else {
			digits[x] = 0
		}

	}
	return append([]int{1}, digits...)

}
func plusOnev2(digits []int) []int {

	chk := false
	n := len(digits)
	for x := n - 1; x >= 0; x-- {
		if digits[x] != 9 {
			digits[x] = digits[x] + 1
			chk = true
			break
		} else {
			digits[x] = 0
		}

	}
	if !chk {
		return append([]int{1}, digits...)
	}

	return digits
}
