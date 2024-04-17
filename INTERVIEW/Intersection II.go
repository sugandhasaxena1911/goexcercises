package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersectionII(nums1, nums2))

}

// non sorted
func intersectionII(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, 0)
	for x := 0; x < len(nums1); x++ {
		v := nums1[x]

		m1[v]++

	}
	result := []int{}
	for x := 0; x < len(nums2); x++ {
		v := nums2[x]
		if _, ok := m1[v]; ok {
			result = append(result, v)
			//delete(m1, v)
			m1[v]--
			if m1[v] == 0 {
				delete(m1, v)
			}
		}

	}
	return result
}
