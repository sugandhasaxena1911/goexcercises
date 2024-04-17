package main

import "fmt"

func main() {
	pascals := generate(5)
	fmt.Println(pascals)

}

func generate(numRows int) [][]int {
	pascals := [][]int{}
	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	if numRows == 2 {
		return [][]int{[]int{1}, []int{1, 1}}
	}
	pascals = append(pascals, []int{1}, []int{1, 1})

	for x := 2; x < numRows; x++ {

		pascalrow := []int{1}

		prevrow := pascals[x-1]
		for y := 0; y < len(prevrow)-1; y++ {
			val := prevrow[y] + prevrow[y+1]
			pascalrow = append(pascalrow, val)
		}
		pascalrow = append(pascalrow, 1)
		pascals = append(pascals, pascalrow)

	}
	return pascals
}
