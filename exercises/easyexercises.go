// Package exercises these are easy exercises
package exercises

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// DivNum tells numbers divided by 7 and not by 5
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

// Vector26 returns a 26 length string vector
func Vector26(str string) string {
	var idx int
	sl := make([]int, 26)
	for _, v := range str {
		idx = int(v - 97) //ASCII for small a
		sl[idx] = sl[idx] + 1
	}
	//sl :26 vector is ready
	//convert int slice to string slice
	sl2 := make([]string, 26)
	for i, v := range sl {
		sl2[i] = strconv.Itoa(v)
	}
	//Slice to string
	return strings.Join(sl2, "#")

}

// takes a slice of strings and returns in the map with sets
func TitlesSet(sl []string) map[string][]string {
	sl2 := make([]string, len(sl))
	map1 := make(map[string][]string)
	var str string
	for _, v := range sl {
		str = Vector26(v)
		fmt.Println("The 26 vector is ", str)
		sl2 = append(map1[str], v)
		map1[str] = sl2

	}
	return map1
}

func FactorialNum(n int) int {
	if n == 1 {
		return 1
	}
	return n * FactorialNum(n-1)

}

// finds the fibonacci upto nth turn  ,nth fibonacci number and return it
func Fibonaccinum(n int) ([]int, int) {
	x, y := 0, 1
	if n == 1 {
		return []int{0}, 0
	}
	if n == 2 {
		return []int{0, 1}, 1
	}
	sl := []int{0, 1}
	num := 0
	for i := 3; i <= n; i++ {
		num = x + y
		sl = append(sl, x+y)
		x, y = y, x+y
	}

	return sl, num
}

func FibNumRecursion(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return FibNumRecursion(n-1) + FibNumRecursion(n-2)
}

// PlusMinus sum of positives , negatives and zeroes
func PlusMinus(arr []int) {
	l := len(arr)
	var p, n, z float64
	for _, v := range arr {
		if v >= 0 {
			if v == 0 {
				z++
			} else {
				p++
			}
		} else {
			n++
		}
	}
	//fmt.Println(l, p, n, z)
	fmt.Printf("%.6f\n", p/float64(l))
	fmt.Printf("%.6f\n", n/float64(l))
	fmt.Printf("%.6f\n", z/float64(l))

}

// minimum and maximum sum of n-1 numbers
func MiniMaxSum(arr []int) {
	var n int
	for _, v := range arr {
		n = n + v
	}

	sort.Ints(arr)
	a := n - arr[0]
	b := n - arr[len(arr)-1]

	fmt.Println(n, arr[0], arr[len(arr)-1])
	fmt.Println(a, b)
}

func TimeConversion(s string) string {
	//hh:mm:ssAM
	var finalTime string

	sub := s[len(s)-2:]
	fmt.Println(sub)
	switch sub {
	case "AM":
		sub2 := s[0:2]
		if sub2 == "12" {
			sub2 = "00"
		}
		finalTime = sub2 + s[2:len(s)-2]
	case "PM":
		sub2 := s[0:2]
		n, e := strconv.Atoi(sub2)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(n)
		if n == 12 {
			n = 12
		} else {
			n = n + 12

		}
		finalTime = strconv.Itoa(n) + s[2:len(s)-2]

	}
	return finalTime

}

// Linear search
func LinearSearch(datalist []int, target int) int {
	for i, v := range datalist {
		if v == target {
			return i
		}
	}
	return -1
}

// Binary search
func BinarySearch(datalist []int, target int) int {
	low := 0
	high := len(datalist) - 1
	for low <= high {
		fmt.Println(low, high)
		mid := (low + high) / 2
		fmt.Println(low, high, mid)
		if datalist[mid] == target {
			return mid
		} else if target > datalist[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}

//Bubblesort

// heres we keep on swapping adjacent element till the largest element reaches the end of array. next loop 0-(n-1)
func BubbleSort(data []int) []int {
	sorted := false
	n := len(data)
	for !sorted {
		swapped := false
		for x := 0; x < n-1; x++ {

			if data[x] > data[x+1] {
				//swapped
				data[x], data[x+1] = data[x+1], data[x]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1 // PLEASE NOTE
	}
	return data
}

// This will return a pivot index and elements on left are less than pivot and on right are greater than pivot
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	pointer := low
	// move elements less than pivot value at start of array, finally move pivot value
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[pointer], arr[j] = arr[j], arr[pointer]
			pointer++
		}
	}
	//finally move pivot at the middle
	arr[pointer], arr[high] = arr[high], arr[pointer]
	return arr, pointer
}
func QuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high) //returns pivot index
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}

// Lonely Integer : there is only one integer thaat occurs once rest all occurs two times
func LonelyInteger(data []int) int {
	check := 0
	m1 := make(map[int]int)
	for _, v := range data {
		m1[v]++
		if m1[v] == 1 {
			check = check + v
		} else {
			check = check - v
		}

	}
	return check
}

func LonelyIntegerV2(data []int) int {
	sort.Ints(data)
	fmt.Println(data)
	n := len(data)
	pt := 0
	var r int
	for x := 0; x < n-1; x++ {
		fmt.Println(data[x], data[x+1])
		if data[x] == data[x+1] {
			pt = 1 // found pair
		} else {
			if pt == 0 {
				r = data[x]
				return r
			}
			pt = 0 // found different number

		}
	}
	if pt == 0 {
		r = data[n-1]
		return r
	}
	return -1
}

func LonelyIntegerV3(data []int) int {
	j := 1
	for j < len(data)-1 {
		if data[j] == data[j-1] {
			j = j + 2
		} else {
			return data[j-1]

		}

	}
	return data[j-1]

}

// Lonely integers list:  x is lonely if occurs only once and x-1 and x+1 not present
func LonelyIntegersList(data []int) []int {
	m1 := make(map[int]int)
	var sl []int

	for _, v := range data {
		m1[v]++
	}
	for k, v := range m1 {
		fmt.Println(k, v, "   ", m1[k+1], m1[k-1])

		// ie it occurs only once and value +1 and value -1 is not present
		if v == 1 && m1[k+1] == 0 && m1[k-1] == 0 {
			sl = append(sl, k)
		}
	}
	return sl
}

// Selection sort  O(nlogn)
// This is very similar to bubblesort . Insted on swapping it keeps note on smaller index and at the end swap it
func SelectionSort(data []int) []int {
	for i, _ := range data {
		min := i
		for x := i; x < len(data); x++ {
			if data[x] < data[min] {
				min = x
			}
		}
		data[i], data[min] = data[min], data[i]
	}
	return data
}

func InsertionSort(data []int) []int {

	for x := 1; x < len(data); x++ {
		key := data[x] //element to be put into sorted array on left
		y := x - 1
		for y = y; y >= 0; y-- {

			if key < data[y] {
				data[y+1] = data[y]
			} else {
				break
			}
		}
		data[y+1] = key

	}
	return data
}

// divides data into multiple mini slices  and then combines
func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}

	}
	for i < len(left) {
		final = append(final, left[i])
		i++
	}
	for j < len(right) {
		final = append(final, right[j])
		j++
	}
	return final
}

// ArmstrongNumber return armstrong numbers till n : sum of cube of digits = number itself
func ArmstrongNumber(n int) []int {
	var result []int
	for x := 1; x <= n; x++ {
		sum := 0
		num := x
		for num > 0 {
			lg := num % 10
			cube := lg * lg * lg
			sum = sum + cube
			num = num / 10
		}
		if sum == x {
			result = append(result, x)
		}

	}
	return result
}

// Print Floyds traingle
func FloydTriangle(rows int) {
	var num int
	for x := 1; x <= rows; x++ {
		for y := 1; y <= x; y++ {
			num++
			fmt.Print(num)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
func SortedExamples() {
	num := []int{50, 90, 30, 10, 50}
	fmt.Println("Int slice before", num)
	/*	Check array of integer Is sorted 	*/
	if sort.IntsAreSorted(num) == false {
		sort.Ints(num) // Sort Integer Array
	}
	fmt.Println("Int slice after", num)
	fmt.Println("search 30 ", sort.SearchInts(num, 30))
	fmt.Println("search 45 ", sort.SearchInts(num, 45)) // NEGATIVE CASE
	index := sort.SearchInts(num, 45)
	if index < len(num) && num[index] == 45 {
		fmt.Println(" FOUND at ", index)
	} else {
		fmt.Println("NOT FOUND. Can be inserted at ", index)
	}

	num2 := []int{50, 90, 30, 10, 50}
	fmt.Println(sort.SearchInts(num2, 45))

	text := []string{"US", "UK", "Germany", "Australia", "Japan"}
	fmt.Println("String slice before", text)
	if sort.StringsAreSorted(text) == false {
		sort.Strings(text) // Sort String Array
	}
	fmt.Println("String slice after", text)
	fmt.Println("Search Japan", sort.SearchStrings(text, "Japan"))
	fmt.Println("Search XYX", sort.SearchStrings(text, "XYX"))

	val := []float64{2.5, 6.3, 1.5, 9.8, 4.7, 1.1, 5.2}
	fmt.Println("Float slice before", val)
	if sort.Float64sAreSorted(val) == false {
		sort.Float64s(val) // Sort Float Array
	}
	fmt.Println("Float slice after", val)
	fmt.Println(sort.SearchFloat64s(val, 5.2))
	fmt.Println("2.1   ", sort.SearchFloat64s(val, 2.1))

}

func PyramidNumbers(n int) {
	for x := 1; x <= n; x++ {
		for y := 1; y <= n-x; y++ {
			fmt.Print(" ")
		}
		num := x*2 - 1
		no := x
		for z := 1; z <= num; z++ {
			fmt.Print(no)
			if z < x {
				no = no + 1
			} else {
				no = no - 1
			}

		}
		fmt.Println("")
	}
}

// Reverse a number
func Reverse(n int) int {
	x := n
	if x < 0 {
		x = x - (x * 2)
	}
	var num int
	for x > 0 {
		lg := x % 10
		num = (num * 10) + lg
		x = x / 10

	}
	if n < 0 {
		num = num - (num * 2)
	}
	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}
	return num

}

// Palindrome check if number is pallindrome
func Palindrome(n int) bool {
	m := Reverse(n)
	if m == n {
		return true
	} else {
		return false
	}
}

func HCF(n1, n2, min int) int {

	if n1%min == 0 && n2%min == 0 {
		return min
	} else {
		return HCF(n1, n2, min-1)
	}

}

func LCM(n1, n2, max int) int {

	if max%n1 == 0 && max%n2 == 0 {
		return max
	} else {
		return LCM(n1, n2, max+1)
	}

}

// RemoveDuplicates from a sorted list : returns unique elemt list  and number of Unique elemnts
func RemoveDuplicates(data []int) ([]int, int) {
	pt := 1
	for x := 1; x < len(data); x++ {
		if data[x] != data[x-1] {
			data[pt] = data[x]
			pt++
		}

	}
	return data[0:pt], pt
}

// MaxProfit profits + sum of max-min

func MaxProfitV2(data []int) int {
	profit := 0
	x := 0
	for x < len(data)-1 {

		for x < len(data)-1 && data[x+1] < data[x] {
			x++
		}
		buy := data[x]
		for x < len(data)-1 && data[x+1] > data[x] {
			x++
		}
		sell := data[x]

		profit = profit + (sell - buy)
		fmt.Println(buy, sell, profit)

	}
	return profit
}

// Rotate k places towards right
func Rotate(data []int, k int) []int {
	fmt.Println(data)
	if k%len(data) == 0 {
		return data
	}
	if k > len(data) {
		k = k % len(data)
	}

	n := len(data)
	data = ReverseSlice(data, 0, n-1)
	data = ReverseSlice(data, 0, k-1)
	data = ReverseSlice(data, k, n-1)
	return data

}

func ReverseSlice(data []int, low int, high int) []int {
	for low < high {
		data[low], data[high] = data[high], data[low]
		low++
		high--
	}
	return data
}
func RotateV2(data []int, k int) []int {
	if k%len(data) == 0 {
		return data
	}
	if k > len(data) {
		k = k % len(data)
	}

	n := len(data)
	fmt.Println(data[n-k:], data[:n-k])
	data = append(data[n-k:], data[:n-k]...)
	return data
}

// IntersectionV2 Returns the common elemnst from the arrays . final array has unique elements

func Intersection(data1, data2 []int) []int {
	m := map[int]int{}
	var result []int
	for _, v := range data1 {
		m[v]++
	}
	for _, v := range data2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			result = append(result, v)
		}
	}

	return result

}

// IntersectionV2 Returns the common elemnst from the arrays . can return common value as many times its appear in both
func IntersectionV2(data1, data2 []int) []int {
	m := map[int]int{}
	var result []int
	for _, v := range data1 {
		m[v]++
	}
	for _, v := range data2 {
		if _, ok := m[v]; ok {
			fmt.Println(v, m[v])
			if m[v] == 0 {
				delete(m, v)
				continue
			}
			m[v]--
			result = append(result, v)
		}
	}

	return result

}

// assumption that both array are sorted
func IntersectionV3(data1, data2 []int) []int {
	m := len(data1)
	n := len(data2)
	result := []int{}
	i, j := 0, 0
	for i < m && j < n {
		if data1[i] < data2[j] {
			i++
		} else if data2[j] < data1[i] {
			j++
		} else {
			result = append(result, data1[i])
			i++
			j++
		}

	}
	//
	/*
		for j < n {
			result = append(result, data2[j])
			j++

		}
		for i < n {
			result = append(result, data1[i])
			i++

		}
	*/

	return result
}

// accepts number disgits in slice and returns number+ 1 in slice
func NumberDigitsInSlice(n []int) []int {

	for x := len(n) - 1; x >= 0; x-- {
		if n[x] < 9 {
			n[x]++
			return n
		}
		n[x] = 0
	}
	m := make([]int, len(n)+1)
	m[0] = 1
	return m
}

func MoveZeroes(n []int) []int {
	//pt := 1
	pointer := 0
	for x := 0; x < len(n); x++ {

		if n[x] != 0 {

			n[pointer] = n[x]
			pointer++
		}
	}

	for pointer < len(n) {
		n[pointer] = 0
		pointer++
	}
	return n
}

var brackets []string

func push(element string) {
	brackets = append(brackets, element)
	fmt.Println("brackets after Push ", brackets)
}

func pop() string {
	if len(brackets) == 0 {
		return " "
	}
	pop := brackets[len(brackets)-1]
	brackets = brackets[:len(brackets)-1]
	fmt.Println("brackets after Pop ", brackets)
	return pop
}

func ValidExpression(st string) bool {
	var s string
	var popped string
	var validity bool
	validity = true
	fmt.Println("validity ", validity)
	for _, v := range st {
		s = string(v)
		fmt.Println("string encountered", s)
		if s == "{" || s == "[" || s == "(" {
			push(s)
		} else {
			if s == "}" || s == "]" || s == ")" {
				popped = pop()
				fmt.Println("popped item ", popped)
				if popped == " " {
					fmt.Println("empty slice , nothing to pop hence invald")
					validity = false
				}
				fmt.Println("popped case check ", popped, s)
				switch popped {
				case "{":
					if s != "}" {
						fmt.Println("hmm")
						validity = false
					}
				case "[":
					if s != "]" {
						validity = false
					}
				case "(":
					if s != ")" {
						validity = false
					}
				}
				if validity == false {
					fmt.Println("please break, dont check more  ")
					break
				}
			}
		}
	}
	if len(brackets) != 0 {
		fmt.Println("len of brackets left to be popped ")
		validity = false
	}
	fmt.Println(validity)
	return validity
}

func ReverseString(st string) string {
	rns := []rune(st)
	fmt.Println("[]rune: ", rns)
	low := 0
	high := len(rns) - 1
	for low < high {
		rns[low], rns[high] = rns[high], rns[low]
		low++
		high--
	}

	return string(rns)
}

func ReverseStringV2(st string) string {
	var str string
	for _, v := range st {
		str = string(v) + str
	}
	return str

}

func ContainsDuplicate(data []int) bool {
	m1 := map[int]int{}
	for i, v := range data {
		if _, ok := m1[v]; ok {
			fmt.Println(i, v)
			return true
		} else {
			m1[v]++

		}
	}
	return false
}

func TwoSumBrute(data []int, target int) (int, int) {
	i := 0
	n := len(data)
	for i < n {
		if data[i] <= target {
			for j := i + 1; j < n; j++ {
				if data[j]+data[i] == target {
					return i, j
				}
			}
			i++
		}
	}
	return -1, -1
}

func TwoSum(data []int, target int) []int {
	i := 0
	n := len(data)
	map1 := map[int]int{}
	for i < n {
		s := target - data[i]
		if v, ok := map1[s]; ok {
			return []int{i, v}
		} else {
			map1[data[i]] = i
		}
		i++
	}

	return []int{-1, -1}
}

func RotateImage(data [][]int) [][]int {
	n := len(data)
	//Transpose the matrix
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			data[i][j], data[j][i] = data[j][i], data[i][j]
		}
	}
	// exchange columns
	for k := 0; k < n; k++ {
		x := 0
		for x < n/2 {
			data[k][x], data[k][n-x-1] = data[k][n-x-1], data[k][x]
			x++
		}
	}
	return data
}

func PrintData(data [][]int) {
	n := len(data)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(data[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func FirstUniqueChar(st string) int {
	m1 := map[string]int{}
	for _, v := range st {
		m1[string(v)]++
	}

	for i, v := range st {
		if m1[string(v)] == 1 {
			return i
		}
	}

	return -1

}

func MergeIntervals(data [][]int) [][]int {
	n := len(data)
	prev := data[0]
	if n == 1 {
		return data
	}
	merge := [][]int{}
	for x := 1; x < n; x++ {
		current := data[x]
		if prev[1] == current[0] {
			node := []int{prev[0], current[1]}
			prev = node

		} else {
			merge = append(merge, prev)
			prev = current

		}

	}
	merge = append(merge, prev)

	return merge

}

func Anagrams(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	map1 := map[string]int{}
	// work if all english characters
	for i, v := range s1 {

		map1[string(v)]++
		map1[string(s2[i])]--
	}

	for _, val := range map1 {

		if val != 0 {
			return false

		}
	}
	return true
}
func AnagramsV2(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	map1 := map[string]int{}
	// work for all  characters
	for _, v := range s1 {
		map1[string(v)]++
	}
	for _, v := range s2 {
		map1[string(v)]--
	}

	for _, val := range map1 {

		if val != 0 {
			return false

		}
	}
	return true
}

func Pallindrome(s string) bool {

	s = RemoveAlphanumeric(s)
	fmt.Println("after removal", s)

	s = strings.ToLower(s)

	fmt.Println("After cap to lower", s)
	low := 0
	high := len(s) - 1
	for low < high {
		fmt.Println(low, high, s[low], s[high])
		if s[low] != s[high] {
			return false
		}
		low++
		high--
		fmt.Println(low, high)
	}

	return true
}

func RemoveAlphanumeric(s string) string {

	var ex = regexp.MustCompile(`[^A-Za-z0-9]+`)
	fmt.Println(ex)
	return ex.ReplaceAllString(s, "")
}

func MyAtoi(s string) int {
	//1. remove whitespaces from start
	s = strings.TrimLeft(s, " ")
	fmt.Println("After whitespace ", s)
	// next character for sign
	if len(s) == 0 {
		return 0
	}
	pt := 0

	signMultiplier := 1
	if s[pt] == '-' {
		signMultiplier = -1
		pt++
	} else if s[pt] == '+' {
		signMultiplier = 1
		pt++

	}
	fmt.Println("sign multiplier ", signMultiplier)

	result := 0
	for pt < len(s) && Isdigit(s[pt]) {
		fmt.Println("INSIDE")
		// ascii 48-57 of 0-9
		n := int(s[pt] - 48)
		fmt.Println(s[pt], n)
		result = (result * 10) + n
		fmt.Println("result", result)

		if signMultiplier == 1 && result > math.MaxInt32 {
			return math.MaxInt32

		}
		if signMultiplier == -1 && result > math.MaxInt32+1 {
			return math.MinInt32

		}
		pt++
	}

	return signMultiplier * result
}
func Isdigit(s byte) bool {
	if s >= 48 && s <= 57 {
		return true
	} else {
		return false
	}
}

func StrStr(hay string, needle string) int {
	i := 0
	j := len(needle)
	high := len(hay)
	//
	for i <= high-j {
		fmt.Println(hay[i], needle[0])
		for i <= high-j && hay[i] != needle[0] {
			fmt.Println(i, hay[i])
			i++
		}
		if i+j <= high && hay[i:i+j] == needle {
			return i
		}
		i++
	}
	return -1

}
func StrStrV2(hay string, needle string) int {

	for x := 0; x <= len(hay)-len(needle); x++ {
		if hay[x:x+len(needle)] == needle {
			return x
		}

	}
	return -1
}
func LongestCommonPrefix(str []string) string {
	pt := 1
	sub := ""
	for pt <= len(str[0]) {
		sub = str[0][0:pt]
		chk := true
		for x := 0; x < len(str); x++ {
			if strings.Index(str[x], sub) == 0 {
				continue
			} else {
				chk = false
				break
			}

		}

		if chk == false {
			return sub[0 : len(sub)-1]
		}

		pt++
	}
	return sub
}
func LongestCommonPrefixV2(str []string) string {
	sort.Strings(str)
	fmt.Println("After sort ", str)
	sub := ""
	first := str[0]
	last := str[len(str)-1]
	for x := 0; x < len(first); x++ {
		if first[x] != last[x] {
			return sub
		} else {
			sub = sub + string(first[x])
		}
	}
	return sub
}

func PassByValueOrRef(str string, n int, sl1 []int, sl2 []int, m1 map[int]int, m2 map[int]int, e Emp, i interface{}, v2 ...int) {
	fmt.Println("String before ", str)
	fmt.Println("int before ", n)
	fmt.Println("slice1  before ", sl1)
	fmt.Println("slice2 before ", sl2)
	fmt.Println("map1 before ", m1)
	fmt.Println("map2 before ", m2)
	fmt.Println("Struct before ", e)
	fmt.Println("Interface before ", i)
	fmt.Println("variadic before ", v2)
	fmt.Println("----------------------------------------\n ")

	str = "String Changed"
	n = 999
	sl1[0] = 999
	sl2 = append(sl2, 999)
	m1[999] = 999
	m2[1] = 888
	e.Name = "String Chnaged"
	i = "Changed string"
	//v2[1] = 999
	v2 = append(v2, 10000)

}
