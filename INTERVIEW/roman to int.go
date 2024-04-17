package main

import (
	"fmt"
)

func main() {
	s := "VX"
	fmt.Println(romanToInt(s))

}

// Will return any value even if any roman is invalid .
// however question says we dont send invalid romans
func romanToInt(s string) int {
	m1 := romansymbols()
	sum := 0

	for i, v := range s {
		sum += m1[string(v)]
		fmt.Println(sum, "BEFORE")
		if i != 0 {
			if m1[string(s[i-1])] < m1[string(v)] {
				sum -= 2 * m1[string(s[i-1])]
			}
			fmt.Println(sum, "AFTER")

		}

	}

	return sum
}

func romansymbols() map[string]int {
	m1 := make(map[string]int)
	m1["I"] = 1
	m1["V"] = 5
	m1["X"] = 10
	m1["L"] = 50
	m1["C"] = 100
	m1["D"] = 500
	m1["M"] = 1000
	return m1

}

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.

*/
