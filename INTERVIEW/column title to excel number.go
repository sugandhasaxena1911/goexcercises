package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(titleToNumber("BY"))
}

/*
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
*/
// 26 number system   :   A*26^0 + A*26^1 = 1+ 26 = 27
// AB = B*26^0 + A*26^1 =  2+ 26= 28
//AZ = 26+ 26= 52
//BX X*26^0 + B*26^1 = 24 +52= 76
func titleToNumber(columnTitle string) int {
	y := 0
	n := 26
	num := 0
	for x := len(columnTitle) - 1; x >= 0; x-- {

		st := (columnTitle[x]) - 64
		num = num + int(st)*int((math.Pow(float64(n), float64(y))))
		y++

	}
	return num
}
