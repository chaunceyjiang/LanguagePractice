package main

import "fmt"

type MinStack struct {
	Head *Node
}
type Node struct {
	Next  *Node
	Value int
	Min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	node := new(Node)
	node.Value = x
	node.Min = x
	node.Next = this.Head
	if this.Head != nil {
		if x < this.Head.Min {
			node.Min = x
		} else {
			node.Min = this.Head.Min
		}
	}

	this.Head = node
}

func (this *MinStack) Pop() {
	this.Head = this.Head.Next
}

func (this *MinStack) Top() int {
	return this.Head.Value
}

func (this *MinStack) GetMin() int {
	return this.Head.Min
}

func main() {
	p := Constructor()
	p.Push(-2)
	p.Push(0)
	p.Push(-3)
	fmt.Println(p.GetMin())
	p.Pop()
	fmt.Println(p.Top())
	fmt.Println(p.GetMin())

}
