package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 3, 5, 6}
	target := 2

	fmt.Println(searchInsertlogn(nums, target))
	fmt.Println(searchInsertGO(nums, target))

}

func searchInsertGO(nums []int, target int) int {

	return sort.SearchInts(nums, target)
}

func searchInsertlogn(nums []int, target int) int {

	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}
	}
	return low
}

// O(n)
func searchInsert(nums []int, target int) int {

	for x := 0; x < len(nums); x++ {
		fmt.Println("inside")
		if nums[x] >= target {
			fmt.Println("inside if")

			return x

		}

	}
	return len(nums)

}
