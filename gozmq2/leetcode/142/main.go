package main



// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast,slow :=head,head
	for fast !=nil &&  fast.Next!=nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow{
			break
		}
	}
	if fast == nil || fast.Next == nil{
		return nil
	}
	slow = head
	for slow != fast{
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}