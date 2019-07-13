package main

import "fmt"

func sortArray(nums []int) []int {
	head,tail:=0,len(nums)-1
	quickSort(head,tail,nums)

	return nums
}
func quickSort(head ,tail int,nums []int) {
	if head >= tail{
		return
	}
	tmp:=nums[head]
	tmp_head:=head
	tmp_tail:=tail
	for head<tail{
		for head<tail && nums[tail]>tmp{
			tail--
		}
		for head<tail && nums[head]<=tmp{
			head++
		}
		swap(head,tail,nums)
	}
	swap(tmp_head,head,nums)
	quickSort(tmp_head,head-1,nums)
	quickSort(head+1,tmp_tail,nums)
}
func swap(a, b int ,nums []int)  {
	t:=nums[a]
	nums[a]=nums[b]
	nums[b]=t
}

func main() {
	fmt.Println(sortArray([]int{5,1,1,2,-1,0,7,6,3,8,11,-5,-7,2,10}))
}