package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index += 1
			nums[index] = nums[i]
		}
	}
	return index + 1
}
func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
