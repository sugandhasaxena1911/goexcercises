package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "Hello,世界"
	/*


		%s	the uninterpreted bytes of the string or slice
		%q	a double-quoted string safely escaped with Go syntax
		%x	base 16, lower-case, two characters per byte
		%X	base 16, upper-case, two characters per byte
	*/
	fmt.Println("length of a= Number of bytes ", len(a))
	//HOW ?????????????
	//answer below
	fmt.Printf("%s %q %x %X\n", a, a, a, a)
	fmt.Println("\n-----------------------------------------------------------------")

	sum := 0
	for i, v := range a {
		// i index of byte , v = value of rune to represent char : UNICODE , a[i]= value of byte at pos i
		fmt.Println("Char : ", string(v))
		fmt.Printf("TYPE Index: %T\n", i)
		fmt.Println("Index: ", i)
		fmt.Printf("TYPE Rune: %T\n", v)
		fmt.Println("Unicode or rune: ", v)
		fmt.Printf("TYPE a[i] : %T\n", a[i])
		fmt.Println("a[i]: first byte of rune OR BYTe at pos i: ", a[i])
		fmt.Println("\n-----------------------------------------------------------------")

		sum++

	}
	fmt.Println("Actual no of characters = Runes count in string = utf8.RuneCountInString  ", sum, utf8.RuneCountInString(a))

	//String si nothing but slice of bytes
	p := []byte(a)
	fmt.Println("Slice of Bytes", p)
	for i, v := range p {
		fmt.Println("INDEX , BYTE ", i, v)
		fmt.Printf("HEX %X\n", v)

	}

	//another example    : array variable cannot be assigned a slice vlaue
	var x int
	arr := [3]int{3, 5, 2}
	//x, arr = arr[0], arr[1:]
	fmt.Println(x, arr)

	/// another example
	arr2 := [...]int{2, 3}
	fmt.Printf("%T", arr2)

	var varX int
	fmt.Println("bhakk", varX)
	fmt.Printf("%T type", varX)

	ax := 0
	if ax != varX {
		fmt.Println("Hello")
	} else {
		fmt.Println("Gello gain")
		fmt.Println(varX)

	}
}
