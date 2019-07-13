package main

import "fmt"

// Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func sortList(head *ListNode) *ListNode {
	var fast,slow ,t *ListNode
	if head == nil{
		return nil
	}
	if head.Next== nil{
		return head
	}
	fast = head
	slow = head
	t = head

	for fast!=nil && fast.Next!=nil {
		fast = fast.Next.Next
		t= slow
		slow = slow.Next

	}
	t.Next = nil

	return migre(sortList(head),sortList(slow))
}

func migre(l1,l2*ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val{
		head = l1
		l1.Next = migre(l1.Next,l2)
	}else{
		head = l2
		l2.Next = migre(l1,l2.Next)
	}

	return head

}

func createList(nums []int) *ListNode {
	var head,p *ListNode

	for i:=0;i<len(nums);i++{
		t:=new(ListNode)
		t.Val = nums[i]
		if p == nil{
			p = t
			head = p
		}else{
			p.Next = t
			p = p.Next
		}
	}
	return head
}

func printList(root *ListNode)  {
	if root == nil{
		return
	}
	fmt.Printf("%d ",root.Val)
	printList(root.Next)
}
func main() {
	l:=createList([]int{1,3,2,7,0,4,5})
	t:=sortList(l)
	printList(t)
}