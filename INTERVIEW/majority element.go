package main

func main() {

}
func majorityElement(nums []int) int {
	m1 := make(map[int]int)
	result := 0
	for x := 0; x < len(nums); x++ {
		m1[nums[x]]++
		if m1[nums[x]] > len(nums)/2 {
			result = nums[x]
		}
	}
	return result
}

// linear time o(n) space O(1)
func majorityElementv2(nums []int) int {
	count := 0
	result := 0
	for x := 0; x < len(nums); x++ {
		if count == 0 {
			result = nums[x]
		}
		if result == nums[x] {
			count++
		} else {
			count--
		}
	}
	return result
}
