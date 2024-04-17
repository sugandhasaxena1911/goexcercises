package main

import "fmt"

func main() {
	haystack := "leetcode"
	needle := "codo"

	fmt.Println(strStr(haystack, needle))

}

func strStr(haystack string, needle string) int {
	n := len(needle)

	for x := 0; x <= len(haystack)-n; x++ {
		if haystack[x:x+n] == needle {
			return x
		}

	}

	return -1
}
