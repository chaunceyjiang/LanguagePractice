package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	Al, Bl := 0, 0
	var res *ListNode
	getListLen(headA, &Al)
	getListLen(headB, &Bl)
	movePoint(headA, headB, Al, Bl, &res)
	return res

}
func movePoint(headA, headB *ListNode, Al, Bl int, res **ListNode) {
	if headB == nil || headA == nil {
		return
	}
	if Al > Bl {
		movePoint(headA.Next, headB, Al-1, Bl, res)
	} else if Al < Bl {
		movePoint(headA, headB.Next, Al, Bl-1, res)
	} else {
		if headA == headB {
			*res = headB
			return
		} else {
			movePoint(headA.Next, headB.Next, Al-1, Bl-1, res)
		}

	}
}
func getListLen(root *ListNode, l *int) {

	if root == nil {
		return
	}
	*l = *l + 1
	getListLen(root.Next, l)
}

func createList(x []int) *ListNode {
	var res, p *ListNode
	for i, v := range x {
		tmp := new(ListNode)
		tmp.Val = v
		if i == 0 {
			res = tmp
		}
		if p == nil {
			p = tmp
		} else {
			p.Next = tmp
			p = p.Next
		}
	}
	return res
}

func main() {
	headA := createList([]int{4, 1, 8, 4, 5})
	headB := createList([]int{5, 0, 1, 8, 4, 5})
	headC := getIntersectionNode(headA, headB)
	fmt.Println(headC.Val)
}
