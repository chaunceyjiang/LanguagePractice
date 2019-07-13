package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
type Deque struct {
	Head *Node
	Tail *Node
}

func (this *Deque) Back() int {
	return this.Tail.Value
}
func (this *Deque) Front() int {
	return this.Head.Value
}
func (this *Deque) PopFront() {
	if this.Head != nil {
		this.Head = this.Head.Next

	}
	if this.Head == nil {
		this.Tail = nil
	} else {
		this.Head.Prev = nil
	}
}
func (this *Deque) PopBack() {
	if this.Tail != nil {
		this.Tail = this.Tail.Prev
	}
	if this.Tail == nil {
		this.Head = nil
	} else {
		this.Tail.Next = nil
	}
}
func (this *Deque) Empty() bool {
	if this.Head == nil && this.Tail == nil {
		return true
	}
	return false

}
func (this *Deque) PushBack(x int) {
	node := new(Node)
	node.Value = x
	if this.Head == nil {
		this.Head = node
	}
	if this.Tail == nil {
		this.Tail = node
	} else {
		node.Prev = this.Tail
		this.Tail.Next = node
		this.Tail = node
	}
}
func (this *Deque) PushFront(x int) {
	node := new(Node)
	node.Value = x
	if this.Head == nil {
		this.Head = node
	} else {
		node.Next = this.Head
		this.Head.Next = node
		this.Head = node
	}
	if this.Tail == nil {
		this.Tail = node
	}
}

type MaxDeque struct {
	deque *Deque
}

func (this *MaxDeque) Push(x int) {
	for (!this.deque.Empty()) && this.deque.Back() < x {
		this.deque.PopBack()
	}
	this.deque.PushBack(x)
}
func (this *MaxDeque) Max() int {
	return this.deque.Front()
}
func (this *MaxDeque) Pop(x int) { //x 从窗口移除的时候
	if !this.deque.Empty() && this.deque.Front() == x {
		this.deque.PopFront()
	}
}
func NewMaxDequr() MaxDeque {
	p := new(Deque)
	return MaxDeque{
		deque: p,
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 1 {
		return nums
	}
	q := NewMaxDequr()
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			q.Push(nums[i])
		} else {
			q.Push(nums[i])
			result = append(result, q.Max())
			q.Pop(nums[i-k+1])
		}

	}

	return result
}
func main() {
	var nums = []int{1, 3, 1, 2, 0, 5}
	fmt.Println(maxSlidingWindow(nums, 3))
}
