package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))

}

func removeElement(nums []int, val int) int {

	pt := 0
	for x := 0; x < len(nums); x++ {

		if nums[x] != val {
			nums[pt] = nums[x]
			pt++

		}
	}

	return len(nums[0:pt])

}
