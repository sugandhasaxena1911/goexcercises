package main

import "fmt"

//Given an array of integers nums and an integer target, return indices of the two numbers
//such that they add up to target.

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
	result = twoSum1(nums, target)
	fmt.Println(result)

}

//O(n2)
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for x := 0; x <= n-2; x++ {
		for y := x + 1; y <= n-1; y++ {
			if nums[x]+nums[y] == target {
				return []int{x, y}
			}
		}

	}

	return []int{0, 0}
}

//O(n)
func twoSum(nums []int, target int) []int {
	m1 := make(map[int]int, 0)
	for i, v := range nums {
		n := target - v
		if _, ok := m1[n]; ok {
			return []int{m1[n], i}
		} else {
			m1[v] = i

		}

	}
	return []int{0, 0}
}
