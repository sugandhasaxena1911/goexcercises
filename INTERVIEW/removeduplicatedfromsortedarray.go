package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)

}

// 11 2 334 556    nums[1]=2 013
func removeDuplicates(nums []int) int {
	pt := 1

	for x := 1; x < len(nums); x++ {
		if nums[x] != nums[x-1] {
			nums[pt] = nums[x]
			pt++
		}
	}

	return len(nums[0:pt])
}
