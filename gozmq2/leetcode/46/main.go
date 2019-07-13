package main

import "fmt"
func permute(nums []int) [][]int {
	switch len(nums) {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{nums}
	case 2:
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	l := 1
	for i := range nums {
		l *= (i + 1)
	}
	ps := make([][]int, l)
	var subs [][]int
	index := 0
	buf := make([]int, len(nums))
	for i := range nums {
		copy(buf, nums)
		if i == 0 {
			subs = permute(nums[i+1:])
		} else if i == len(nums)-1 {
			subs = permute(nums[0:i])
		} else {
			subs = permute(append(buf[0:i], buf[i+1:]...))
		}
		for j := range subs {
			ps[index] = append([]int{nums[i]}, subs[j]...)
			index++
		}
	}
	return ps
}
/*func permute(nums []int) [][]int {
	res := make([][]int, 0)
	swap(nums, []int{}, &res)
	return res
}
func swap(path, tmp []int, res *[][]int) {
	if len(path) == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(path); i++ {
		t_path := make([]int, len(path)-1)
		index := 0
		for j := 0; j < len(path); j++ {
			if j == i {
				continue
			}
			t_path[index] = path[j]
			index++
		}
		tmp = append(tmp, path[i])
		swap(t_path, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}

}*/
func main() {
	res := permute([]int{1, 2, 3})
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
