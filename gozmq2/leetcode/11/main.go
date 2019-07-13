package main

import "fmt"

func maxArea(height []int) int {
	var curMaxArea int
	head,tail:=0,len(height)-1
	for head<tail{
		if height[head] >= height[tail] {
			if height[tail] * (tail-head) > curMaxArea{
				curMaxArea = height[tail] * (tail-head)
			}
			tail--
		}else {
			if height[head] * (tail-head) > curMaxArea{
				curMaxArea = height[head] * (tail-head)
			}
			head++
		}
	}
	return curMaxArea
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}