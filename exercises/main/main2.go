package main

import (
	"fmt"
	"github.com/sugandhasaxena1911/goexcercises/exercises"
)

func main() {

	/*

				s := []int{1, 6, 3, -9, -4, 0, 2, -1, 0}
				//exercises.PlusMinus(s)

				exercises.MiniMaxSum(s)
				st := exercises.TimeConversion("12:33:11PM")
				fmt.Println(st)

				fmt.Println("LINEAR SEARCH")
				data := []int{2, 5, 9, 1, 0, -4, 3}
				target := 9
				fmt.Println(exercises.LinearSearch(data, target))
				fmt.Println("BINARY SEARCH")
				sort.Ints(data)
				fmt.Println(data)
				fmt.Println(exercises.BinarySearch(data, target))

				fmt.Println("BUBBLE SORT")
				data = []int{2, 5, -6, 1, 0, -4, 3, 15}
				data1 := []int{7}

				fmt.Println(data)
				fmt.Println("sorted: \n", exercises.BubbleSort(data))
				fmt.Println(data1)
				fmt.Println("sorted: \n", exercises.BubbleSort(data1))

				fmt.Println("QUICKSORT")
				data = []int{2, 5, -6, 1, 0, -4, 3, 15}
				data1 = []int{7}
				//7 == len(arr)-1
				fmt.Println("sorted: \n", exercises.QuickSort(data, 0, 7))

				fmt.Println("LONELY INTEGER")
				data = []int{0, 0, 2, 5, 1, 5, 2, 1, 7}
				fmt.Println("The lonely int is ", exercises.LonelyInteger(data))

				fmt.Println("LONELY INTEGERS LIST ")
				data = []int{0, 0, 8, 5, 1, 5, 3}
				fmt.Println(data, "\nThe lonely integer list is ", exercises.LonelyIntegersList(data))

				fmt.Println("SELECTION SORT")
				data = []int{0, 0, 8, 5, 1, 5, 3, -6}
				fmt.Println(data)
				fmt.Println("sorted array ", exercises.SelectionSort(data))

				fmt.Println("INSERTION SORT")
				data = []int{0, 0, 8, 5, 1, 5, 3, -6, 4, -9, 7}
				fmt.Println(data)
				fmt.Println("sorted array ", exercises.SelectionSort(data))

				fmt.Println("MERGE SORT")
				data = []int{0, 0, 8, 5, 1, 5, 3, -6, 4, -9, 7, 2, -2}
				fmt.Println(data)
				fmt.Println("sorted array ", exercises.MergeSort(data))

				fmt.Println("ARMSTRONG")
				sl := exercises.ArmstrongNumber(500)
				fmt.Println("values :   ", sl)

				fmt.Println("FLOYDS TRIANGLE")
				exercises.FloydTriangle(10)

				fmt.Println("SORTED")
				exercises.SortedExamples()

				fmt.Println("PYRAMID")
				exercises.PyramidNumbers(5)

				fmt.Println("REVERSE NUMBER")
				n := 579
				fmt.Println("Reverse of ", n, " is :", exercises.Reverse(n))

				fmt.Println("PALINDROME")
				n = 5445
				fmt.Println("number is palindrome ????", exercises.Palindrome(n))

				fmt.Println("HCF")
				n1, n2 := 40, 16
				if n1 <= n2 {
					fmt.Println(exercises.HCF(n1, n2, n1))
				} else {
					fmt.Println(exercises.HCF(n1, n2, n2))

				}

				fmt.Println("LCM")
				n1, n2 = 17, 9
				if n1 >= n2 {
					fmt.Println(exercises.LCM(n1, n2, n1))
				} else {
					fmt.Println(exercises.LCM(n1, n2, n2))

				}

				data = []int{0, 1, 2, 2, 3, 3, 3, 4, 8}
				fmt.Println("REMOVE DUPLICATES")
				fmt.Println(data)
				d, n := exercises.RemoveDuplicates(data)
				fmt.Println(d, n)

				fmt.Println("MAX PROFIT")

				data = []int{5, 7, 9, 10, 3, 2, 4, 8, 5, 10}
				fmt.Println(data)

				fmt.Println(exercises.MaxProfitV2(data))

				data = []int{0, 1, 2, 2, 3, 3, 4, 8}

				fmt.Println(exercises.Rotate(data, 2))
				data = []int{0, 1, 2, 2, 3, 3, 4, 8}

				fmt.Println(exercises.RotateV2(data, 2))

				fmt.Println("INTERSECTION")
				data1 = []int{0, 1, 2, 2, 3, 5, 4, 8, 4}
				data2 := []int{0, 3, 2, 5, 4, 3, 4, 9, 4}
				fmt.Println(exercises.Intersection(data1, data2))

				fmt.Println("INTERSECTIONV2")
				data1 = []int{0, 1, 2, 2, 3, 5, 4, 8, 4}
				data2 = []int{0, 3, 2, 5, 4, 3, 4, 9, 4}
				fmt.Println(exercises.IntersectionV2(data1, data2))

				fmt.Println("INTERSECTIONV3")
				data1 = []int{0, 1, 2, 2, 3, 5, 4, 8, 4}
				data2 = []int{0, 3, 2, 5, 4, 3, 4, 9, 4}
				sort.Ints(data1)
				sort.Ints(data2)
				fmt.Println(exercises.IntersectionV3(data1, data2))

				fmt.Println("NumberDigitsInSlice")
				fmt.Println(exercises.NumberDigitsInSlice([]int{9, 9, 9}))

				fmt.Println("MOVEZEROES")
				data1 = []int{7, 1, 2, 0, 0, 3, 5, 0, 7, 3, 0, 5, 6, 0}
				fmt.Println(data1)

				fmt.Println("Output", exercises.MoveZeroes(data1))




			d := []int{0, 2, 3, 4, 1, 8, 5, 2, 90}
			fmt.Println("Contains duplicate ", exercises.ContainsDuplicate(d))

			fmt.Println("LONELY INTEGER")
			d1 := []int{0, 9, 4, 9, 0, 3, 4}
			fmt.Println(d1)
			fmt.Println(exercises.LonelyIntegerV2(d1))
			fmt.Println(exercises.LonelyIntegerV3(d1))

			fmt.Println("TWOSUM")
			d2 := []int{2, 7, 11, 15}
			d3 := 13
			fmt.Println(exercises.TwoSumBrute(d2, d3))
			fmt.Println(exercises.TwoSum(d2, d3))

			fmt.Println("RotateImage")
			d4 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
			exercises.PrintData(d4)
			fmt.Println("")
			exercises.PrintData(exercises.RotateImage(d4))

			fmt.Println("Reverse String")
			st := "Hello, 世界"
			fmt.Println(exercises.ReverseString(st))
			fmt.Println(exercises.ReverseStringV2(st))

			fmt.Println("First Unique")
			fmt.Println(exercises.FirstUniqueChar("SugandhaSun"))

			fmt.Println("MERGE Intervals")

			d5 := [][]int{{0, 2}, {2, 5}, {5, 9}, {10, 13}, {13, 14}, {15, 17}}
			fmt.Println(d5)
			fmt.Println(exercises.MergeIntervals(d5))
			fmt.Println("\n")
			d6 := [][]int{{0, 2}, {2, 5}, {5, 9}, {10, 13}, {13, 14}, {14, 17}}
			fmt.Println(d6)
			fmt.Println(exercises.MergeIntervals(d6))

			fmt.Println("Anagrams")
			fmt.Println(exercises.Anagrams("hello界", "界ollhe"))
			fmt.Println(exercises.AnagramsV2("hello界", "界ollhe"))





		fmt.Println(exercises.Pallindrome("A man, a plan, a canal: Panama"))
		fmt.Println(exercises.Pallindrome("race a car"))

		//fmt.Print(strings.TrimLeft("¡¡¡!Hello, Gophers!!!", "!¡"))
		fmt.Println("ATOI")
		fmt.Println("value is ", exercises.MyAtoi("  -2356"))
		fmt.Println("value is ", exercises.MyAtoi("+1"))

	*/

	/*
		fmt.Println("STRSTR")
		//fmt.Println(exercises.StrStr("sadbutsad", "sad"))
		fmt.Println(exercises.StrStr("llleetcodeleeto", "leeto"))

		fmt.Println("LONGEST COMMON PREFIX")
		d10 := []string{"flower", "flop", "flow", "floss"}
		d11 := exercises.LongestCommonPrefix(d10)
		fmt.Println(d11)

		d12 := []string{"floppers", "flop", "flow", "flops"}
		d13 := exercises.LongestCommonPrefixV2(d12)
		fmt.Println(d13)

	*/

	m1 := map[int]int{}
	m1[0] = 0
	m1[1] = 1

	m2 := map[int]int{}
	m2[0] = 100
	m2[1] = 101
	e := exercises.Emp{"sugandha", 1}
	var i interface{}
	i = "hello"
	slc := []int{0, 1, 2, 3, 4}
	str := "Sugandha "
	n := 0
	sl1 := []int{1, 2, 3}
	sl2 := []int{10, 20, 30}

	exercises.PassByValueOrRef(str, n, sl1, sl2, m1, m2, e, i, slc...)
	fmt.Println("String after ", str)
	fmt.Println("int after ", n)
	fmt.Println("slice1  after ", sl1)
	fmt.Println("slice2 after ", sl2)
	fmt.Println("map1 after ", m1)
	fmt.Println("map2 after ", m2)
	fmt.Println("Struct after ", e)
	fmt.Println("Interface after ", i)
	fmt.Println("variadic after ", slc)

	fmt.Println("--------------------------------------------\n")
	fmt.Println("STRING is PASS BY VALUE")
	fmt.Println("INT is PASS BY VALUE")
	fmt.Println("SLICE is PASS BY REFERENCE, NOTE : append func creates another array & hence changes dont reflect")
	fmt.Println("MAP is PASS BY REFERENCE, NOTE : both adding elements & modify reflects changes ")
	fmt.Println("STRUCT is PASS BY VALUE")
	fmt.Println("INTERFACE is PASS BY VALUE")
	fmt.Println("VARIADIC PARAM is PASS BY REFERENCE , Note : just like slices")

}
