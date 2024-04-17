package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	start := 0
	ways := 0
	for start < n {

		start = start + 2
		ways = ways + 2
		if start > n {
			ways--
		}
	}
	return ways
}
