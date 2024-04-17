package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "0P"
	fmt.Println(isPalindromes(s))
}

func isPalindromes(s string) bool {

	// a-z 97 -122
	s = strings.ToLower(s)
	fmt.Println(s)
	snew := ""
	for x := 0; x < len(s); x++ {
		if s[x] >= 97 && s[x] <= 122 {
			snew = snew + string(s[x])
		}
	}
	fmt.Println(snew)

	return checkpallindrome(snew)
}

func checkpallindrome(s string) bool {

	low := 0
	high := len(s) - 1

	for low < high {

		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}

// remove non alpha
func RemoveNOnAlphanumeric(s string) string {

	var ex = regexp.MustCompile(`[^A-Za-z0-9]+`)
	fmt.Println(ex)
	return ex.ReplaceAllString(s, "")
}
