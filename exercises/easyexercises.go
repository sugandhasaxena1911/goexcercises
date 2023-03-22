// Package exercises these are easy exercises
package exercises

import (
	"strconv"
	"strings"
)

// DivNum divides numbers
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
