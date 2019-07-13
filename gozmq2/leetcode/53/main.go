package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := 0
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
