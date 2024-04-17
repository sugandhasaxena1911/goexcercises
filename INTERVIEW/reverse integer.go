package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(reverse(1534236469))

}

func reverse(x int) int {
	n := x

	if n < 0 {

		n = n - (2 * n)
	}

	//123= 3+((3*10)+2) || ((32*10)+1)||   final : 321
	var num int
	for n > 0 {
		lg := n % 10
		num = (num * 10) + lg
		n = n / 10
	}
	if x < 0 {
		num = num - (num * 2)
	}
	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}
	return num
}
