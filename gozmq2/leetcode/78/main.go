package main

import "fmt"

func subsets(nums []int) [][]int {
	res := new([][]int)
	Backtracking(0, []int{}, nums, res)
	*res = append(*res, []int{})
	return *res
}

func Backtracking(i int, tmp []int, nums []int, res *[][]int) {
	for j := i; j < len(nums); j++ {
		tmps := append(tmp, nums[j])
		t := make([]int, len(tmps))
		copy(t, tmps)
		*res = append(*res, t)
		Backtracking(j+1, tmps, nums, res)
	}

}
func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
