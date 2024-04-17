package main

import (
	"fmt"
	"sort"
	"strings"
)

//Write a function to find the longest common prefix string amongst an array of strings.

func main() {
	str := []string{"flower", "flo"}
	st := longestCommonPrefix(str)
	fmt.Println(st)
	st = longestCommonPrefixv2(str)
	fmt.Println(st)

}
func longestCommonPrefix(strs []string) string {
	//chk := true
	idx := 1
	commonprefix := ""
	for idx <= len(strs[0]) {
		commonprefix = strs[0][0:idx]
		for x := 1; x < len(strs); x++ {
			st := strs[x]
			if strings.Index(st, commonprefix) != 0 {
				//chk = false
				//break
				return commonprefix[0 : len(commonprefix)-1]
			}

		}

		/*if chk == false {
			return commonprefix[0 : len(commonprefix)-1]
		}
		*/
		idx++
	}
	return commonprefix
}

func longestCommonPrefixv2(strs []string) string {
	var commonprefix string
	sort.Strings(strs)
	st1 := strs[0]
	st2 := strs[len(strs)-1]
	for x := 0; x < len(st1); x++ {
		commonprefix = st1[0 : x+1]
		if strings.Index(st2, commonprefix) != 0 {
			return commonprefix[0 : len(commonprefix)-1]

		}
	}

	return commonprefix
}
