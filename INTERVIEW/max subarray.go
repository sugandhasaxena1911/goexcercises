package main

func main() {

}

// -2,1,-3,4,-1,2,1,-5,4
func maxSubArray(nums []int) int {
	maxsofar := nums[0]
	currsum := nums[0]
	for x := 1; x < len(nums); x++ {
		val := currsum + nums[x]
		currsum = max(val, nums[x])
		maxsofar = max(currsum, maxsofar)

	}
	return maxsofar
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
