package Arrays

import (
	"fmt"
	"sort"
)

//https://medium.com/javarevisited/20-array-coding-problems-and-questions-from-programming-interviews-869b475b9121

// MissingNUm How do you find the missing number in a given integer array of 1 to 100: no duplicates or more than one missing
func MissingNUm(sl []int) int {
	//calculate the sum from 1 to 100
	n := len(sl) + 1
	sum := n * (n + 1) / 2
	s := 0
	for _, v := range sl {
		s = s + v
	}
	return sum - s
}

// IsDuplicatesPresent Does this array has duplicate/s , return true/false and  number of duplicates  ?
func IsDuplicatesPresent(sl []int) (bool, map[int]int) {
	check := false
	//map2 := make (map[int]int)
	map1 := make(map[int]int)
	for _, v := range sl {
		if _, ok := map1[v]; ok == true {
			check = true
			map1[v] = map1[v] + 1

		} else {
			map1[v] = map1[v] + 1
		}
	}
	return check, map1
}

func MissingNumbers(sl []int, n int) []int {
	_, map1 := IsDuplicatesPresent(sl)
	var sl2 []int
	for x := 1; x <= n; x++ {
		if _, ok := map1[x]; ok == false {
			sl2 = append(sl2, x)
		}
	}
	return sl2
}

func RemoveDuplicates(sl []int) []int {
	map1 := make(map[int]int)
	chk := false
	var sl2 []int
	for _, v := range sl {
		if _, ok := map1[v]; ok == true {
			chk = true
			//duplicate found
			map1[v] = map1[v] + 1
		} else {
			map1[v] = map1[v] + 1
		}

	}
	if chk == true {
		for key, _ := range map1 {

			sl2 = append(sl2, key)

		}
	} else {
		return sl
	}

	return sl2

}

// LargestSmallest largest and smallest number in the list passed
func LargestSmallest(sl []int) (int, int) {
	largest := sl[0]
	smallest := sl[0]
	for _, v := range sl {
		if v > largest {
			largest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	return largest, smallest
}

// PeakNumbers returns the number which is greater than or equal to its neighbour
func PeakNumbers(al []int) []int {
	var sl []int
	for i, v := range al {
		{
			if i == 0 && v > al[i+1] {
				sl = append(sl, v)
			} else if i == len(al)-1 && v > al[i-1] {
				sl = append(sl, v)
			} else {

				if v > al[i-1] && v > al[i+1] {
					sl = append(sl, v)
				}
			}
		}
	}

	return sl
}

func FindPair(sl []int, sum int) [][]int {
	//sort a slice
	sort.Ints(sl)
	fmt.Println(sl)
	var final [][]int
	var pairs []int
	if len(sl) == 2 {
		fmt.Println("hello")

		if sum == sl[0]+sl[1] {
			pairs = append(pairs, sl[0], sl[1])
			final = append(final, pairs)
			return final
		}
	}
	for i, v := range sl {

		sl2 := sl[i+1:]
		for _, v1 := range sl2 {
			if sum == (v + v1) {
				final = append(final, []int{v, v1})

			}
		}
	}

	return final

}

// ReverseSlice This function will swap the values
func ReverseSlice(sl []int) []int {
	sl2 := make([]int, len(sl)) // Note this important point
	copy(sl2, sl)
	for i, j := 0, len(sl2)-1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	return sl2
}

// MoveNegLeftPosRight Move -ve numbers towards left and positive towards right
func MoveNegLeftPosRight(sl []int) []int {
	x := 0
	y := len(sl) - 1
	for x < y {
		fmt.Println(x, y)

		if sl[x] < 0 {
			x++
			fmt.Println("x moved")
		}
		if sl[y] >= 0 {
			y--
			fmt.Println("y moved")

		}

		if sl[x] >= 0 && sl[y] < 0 {
			sl[x], sl[y] = sl[y], sl[x]
			fmt.Println("swapped")
			x++
			y--

		}

	}
	return sl

}

// SortArray012s Dutch National Flag problem
func SortArray012s(sl []int) []int {
	x := 0
	y := len(sl) - 1
	mid := 0
	//0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1
	for mid < y {
		fmt.Println(x, y, mid, sl)
		if sl[mid] == 1 {
			mid++
		}

		if sl[x] == 0 {
			x++
		}
		if sl[y] == 2 {
			y--

		}

		if sl[mid] == 0 && sl[x] == 1 {
			sl[x], sl[mid] = sl[mid], sl[x]
			x++
			mid++
		}
		if sl[mid] == 2 && sl[y] == 1 {
			sl[y], sl[mid] = sl[mid], sl[y]
			y--
			mid++
		}
	}
	return sl

}

// UnionIntersect finds the union and intersect of array
func UnionIntersect(sl1, sl2 []int) ([]int, []int) {
	x, y := 0, 0
	var union []int
	var intersect []int

	//sl1 = RemoveDuplicates(sl1)
	//sl2 = RemoveDuplicates(sl2)
	//fmt.Println("After removing dups ", sl1, sl2)
	sort.Ints(sl1)
	sort.Ints(sl2)
	fmt.Println("After sort ", sl1, sl2)
	i := len(sl1)
	j := len(sl2)
	var prev int

	for x < i && y < j {
		fmt.Println("start", x, y, sl1[x], sl2[y])
		if sl1[x] < sl2[y] {
			if sl1[x] != prev {
				fmt.Println("append x ", sl1[x], prev)

				union = append(union, sl1[x])
			}
			prev = sl1[x]
			x++

		} else if sl2[y] < sl1[x] {
			if sl2[y] != prev {
				fmt.Println("append y ", sl2[y], prev)

				union = append(union, sl2[y])
			}
			prev = sl2[y]
			y++
		} else {
			if sl1[x] != prev {
				fmt.Println("append either  ", sl2[y], prev)

				union = append(union, sl1[x])
				intersect = append(intersect, sl1[x])
			}
			prev = sl1[x]
			x++
			y++
			fmt.Println("append xy  ", x, y)

		}
	}
	for x < i {
		if sl1[x] != prev {
			union = append(union, sl1[x])
		}
		prev = sl1[x]
		x++
	}
	for y < j {
		if sl2[y] != prev {
			union = append(union, sl2[y])

		}
		prev = sl2[y]
		y++
	}

	return union, intersect

}
