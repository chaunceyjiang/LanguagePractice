package main

import "fmt"

func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	res_left := make([]int, len(nums))
	res_left[0] = 1
	res_right := make([]int, len(nums))
	res:=make([]int,len(nums))
	res_right[len(nums)-1] = 1
	for i, j := 0, len(nums)-1; i < len(nums)-1; i, j = i+1, j-1 {
		res_left[i+1] =nums[i] * res_left[i]
		res_right[j-1] = nums[j] * res_right[j]
		if i>=j{
			res[i] = res_left[i] * res_right[i]
			res[j] = res_left[j] * res_right[j]
		}
	}
	res[0] = res_left[0] * res_right[0]
	res[len(nums)-1] = res_left[len(nums)-1] * res_right[len(nums)-1]
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
}