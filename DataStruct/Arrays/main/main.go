package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/DataStruct/Arrays"
)

func main() {
	//Missing num
	sl := []int{10, 4, 2, 5, 9, 6, 1, 7, 3}
	fmt.Println("the missing number is ", Arrays.MissingNUm(sl))

	// duplicates present
	sl1 := []int{2, 4, 6, 7, 8, 0, 4, 2, 7}
	b, map1 := Arrays.IsDuplicatesPresent(sl1)
	fmt.Println("The duplicates are present ", b)
	fmt.Println("The duplicates are  ", map1)

	//largest smallest
	l, s := Arrays.LargestSmallest(sl1)
	fmt.Println("The largest and smallest are ", l, s)

	//missing numbers   , n=10
	sl = []int{10, 4, 2, 1, 10, 9, 6, 1, 6}
	n := 10
	fmt.Println("The missing numbers are ", Arrays.MissingNumbers(sl, n))

	// find pairs
	sl = []int{9, 4, 2, 5, 1, 5, 8, 6}

	pairs := Arrays.FindPair(sl, 10)
	fmt.Println("the pairs are for sum 10", pairs)

	//remove duplicates
	sl = []int{9, 4, 2, 5, 1, 5, 8, 6}
	fmt.Println(sl, "\nafter removing duplicates is ", Arrays.RemoveDuplicates(sl))

	//PeakNUmbers
	fmt.Println("Peaks in ", sl, "are \n", Arrays.PeakNumbers(sl))

	// reverse slice
	slrev := []int{1, 2, 7, 6, 4, 0, 45}
	fmt.Println("reverse of ", slrev, "is : \n", Arrays.ReverseSlice(slrev))

	//MoveNegLeftPosRight
	slmove := []int{12, 11, -13, -5, 6, -7, 5, -3, -6}
	fmt.Println(Arrays.MoveNegLeftPosRight(slmove))
	//Sort012
	//slsort := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	//fmt.Println(Arrays.SortArray012s(slsort))

	// union intersect
	sl1 = []int{2, 5, 6, 8, 9, 10, 6, 2, 1}
	sl2 := []int{1, 2, 3, 6, 9, 11, 11}

	u, i := Arrays.UnionIntersect(sl1, sl2)
	fmt.Println("Union :", u, "\nIntersect : ", i)
}
