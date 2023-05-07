package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/exercises"
)

func main() {
	//DivNum
	finalstr := exercises.DivNum(2000, 3200)
	fmt.Println("The final output ", finalstr)
	//excercise -1
	titles := []string{"manu", "abbccc", "ccabccb", "bcaccb", "tina", "itna", "nati", "maun", "muna"}
	var titlesmap map[string][]string
	titlesmap = exercises.TitlesSet(titles)
	fmt.Println(titlesmap)
	query := "anum"
	st := exercises.Vector26(query)
	fmt.Println("The vector of word is  ", st)
	fmt.Println("The matching words of", query, "are ", titlesmap[st])

	//factorial of num
	num := 7
	fmt.Println("The factorial of", num, "is : \t", exercises.FactorialNum(num))
	//Fibonaccinum
	num = 13
	slfib, n := exercises.Fibonaccinum(num)
	fmt.Println("The fibonacci numbers upto ", num, "th turn are: \n", slfib, "\nThe last fibonacci is number : ", n)

	// recursion
	fmt.Println("The ", num, "th fibonacci number is: ", exercises.FibNumRecursion(num))
}
