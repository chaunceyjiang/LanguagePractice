package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)

	findMixNode(l1, l2, result)
	return result
}
func findMixNode(l1 *ListNode, l2 *ListNode, result *ListNode) {

	tmp := new(ListNode)
	result.Next = tmp
	if l1.Val < l2.Val {
		result.Val = l1.Val
		if l1.Next == nil {
			result.Next = l2
			return
		}
		findMixNode(l1.Next, l2, result.Next)
	} else {
		result.Val = l2.Val
		if l2.Next == nil {
			result.Next = l1
			return
		}
		findMixNode(l1, l2.Next, result.Next)
	}

}

func createList(value []int) *ListNode {

	p := new(ListNode)
	root := p
	for i := 0; i < len(value)-1; i = i + 1 {

		p.Val = value[i]
		p.Next = new(ListNode)
		p = p.Next
	}
	p.Val = value[len(value)-1]
	return root
}
func printList(root *ListNode) {

	if root != nil {
		fmt.Println(root.Val)
		printList(root.Next)
	}

}
func main() {
	v1 := []int{1, 2, 4}
	v2 := []int{1, 3, 4}
	l1 := createList(v1)
	l2 := createList(v2)

	printList(mergeTwoLists(l1, l2))
}
