package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 {
		if len(nums2) != 0 {
			for i := 0; i < len(nums2); i++ {
				nums1[i] = nums2[i]
			}
		}
		return
	}
	if len(nums2) == 0 {
		return
	}
	for i := 0; i < m; i++ {
		if nums1[i] > nums2[0] {
			t := nums1[i]
			nums1[i] = nums2[0]
			nums2[0] = t
			swap(nums2)
		}
	}
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}

}
func swap(nums []int) {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] > nums[i] {
			t := nums[i]
			nums[i] = nums[index]
			nums[index] = t
			index += 1
		}

	}
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
