// Package exercises these are easy exercises
package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

// DivNum tells numbers divided by 7 and not by 5
func DivNum(lower_l int, upper_l int) string {
	var str []string
	for x := lower_l; x <= upper_l; x++ {
		if x%7 == 0 && x%5 != 0 {
			//fmt.Println(strconv.Itoa(x))
			str = append(str, strconv.Itoa(x))
		}
	}
	return strings.Join(str, ",")
}

// Vector26 returns a 26 length string vector
func Vector26(str string) string {
	var idx int
	sl := make([]int, 26)
	for _, v := range str {
		idx = int(v - 97) //ASCII for small a
		sl[idx] = sl[idx] + 1
	}
	//sl :26 vector is ready
	//convert int slice to string slice
	sl2 := make([]string, 26)
	for i, v := range sl {
		sl2[i] = strconv.Itoa(v)
	}
	//Slice to string
	return strings.Join(sl2, "#")

}

// takes a slice of strings and returns in the map with sets
func TitlesSet(sl []string) map[string][]string {
	sl2 := make([]string, len(sl))
	map1 := make(map[string][]string)
	var str string
	for _, v := range sl {
		str = Vector26(v)
		fmt.Println("The 26 vector is ", str)
		sl2 = append(map1[str], v)
		map1[str] = sl2

	}
	return map1
}

func FactorialNum(n int) int {
	if n == 1 {
		return 1
	}
	return n * FactorialNum(n-1)

}

// finds the fibonacci upto nth turn  ,nth fibonacci number and return it
func Fibonaccinum(n int) ([]int, int) {
	x, y := 0, 1
	if n == 1 {
		return []int{0}, 0
	}
	if n == 2 {
		return []int{0, 1}, 1
	}
	sl := []int{0, 1}
	num := 0
	for i := 3; i <= n; i++ {
		num = x + y
		sl = append(sl, x+y)
		x, y = y, x+y
	}

	return sl, num
}

func FibNumRecursion(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return FibNumRecursion(n-1) + FibNumRecursion(n-2)
}
