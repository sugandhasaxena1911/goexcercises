package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindromev2(1221))

}

/*
Given an integer x, return true if x is a
palindrome
, and false otherwise.
*/
func isPalindrome(x int) bool {

	if x == reversenum(x) {
		return true
	} else {
		return false
	}

}

// doessnt work . however first three conds of edge cases are good
func isPalindromev2(x int) bool {
	//1221
	// negative nums cannot be palindrome
	if x < 0 {
		return false
	}
	// single digit cannot be palindrome
	if x < 10 {
		return true
	}
	// if last digit 0 then cannot be palindrome
	if x%10 == 0 {
		return false
	}
	n := x
	for n > 0 {
		lg := n % 10
		fg := n / 10
		if fg != lg {
			return false
		}
		n = n / 10
	}
	return true
}

func reversenum(x int) int {

	n := x
	num := 0
	for n > 0 {
		lg := n % 10
		num = (num * 10) + lg
		n = n / 10
	}
	return num
}
