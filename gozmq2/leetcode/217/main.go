package main

import (
	"fmt"
	"strconv"
)

func containsDuplicate(nums []int) bool {
	m := make(map[string]bool)
	for _, v := range nums {
		vv := strconv.Itoa(v)
		if _, ok := m[vv]; ok {
			return true
		}
		m[vv] = true
	}
	return false
}
func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
