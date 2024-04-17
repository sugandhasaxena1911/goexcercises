package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(lengthOfLastWord(" moon"))

}
func lengthOfLastWordv2(s string) int {
	s = trimspacesright(s)
	fmt.Println(s)
	l := 0
	for x := len(s) - 1; x >= 0; x-- {
		if s[x] != 32 {
			l++
		} else {
			break
		}

	}
	return l
}

func trimspacesright(s string) string {
	// hello
	pt := 0
	for x := len(s) - 1; x >= 0; x-- {
		fmt.Println("inside ", s[x])
		if s[x] != 32 {
			pt = x
			break
		}
	}
	return s[0 : pt+1]
}
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	return len(s) - strings.LastIndex(s, " ") - 1

}
