package main

import (
	"fmt"
)

type Node struct {
	prev, next *Node
	key        int
}
type Value struct {
	value int
	node  *Node
}
type LRUCache struct {
	cache      map[int]*Value
	head, tail *Node
	capacity   int
	cacheCap   int
}

func Constructor(capacity int) LRUCache {
	var lrucache LRUCache
	lrucache = LRUCache{
		cache:    make(map[int]*Value),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		cacheCap: 0,
	}

	return lrucache
}

func (this *LRUCache) Get(key int) int {
	var v *Value
	var ok bool
	var tailNode *Node
	if v, ok = this.cache[key]; !ok {
		return -1
	}
	tailNode = v.node.prev
	if v.node.prev !=nil{
		v.node.prev.next = v.node.next
	}
	if v.node.next !=nil{
		v.node.next.prev = v.node.prev
	}

	this.head.prev = v.node
	v.node.next = this.head
	this.head = v.node
	v.node.prev = nil
	if this.tail == this.head && tailNode!=nil{
		this.tail = tailNode
	}
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	if _,ok:=this.cache[key];!ok{
		this.cacheCap++
		if this.cacheCap > this.capacity {
			delNode := this.tail
			delete(this.cache, delNode.key)
			this.tail = this.tail.prev
		}
	}

	var v *Value
	node := &Node{
		prev: nil,
		next: nil,
		key:  key,
	}
	if this.head == nil {
		this.head = node
	}else {
		node.next=this.head
		this.head.prev=node
		this.head=node
	}
	if this.tail == nil {
		this.tail = node
	}
	v = &Value{
		value: value,
		node:  node,
	}

	this.cache[key] = v
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	var param_1 int
	obj := Constructor(2)
	//param_1 =obj.Get(2)
	//fmt.Println(param_1)
	obj.Put(2, 6)
	obj.Put(1,1)
	param_1 =obj.Get(1)
	fmt.Println(param_1)

	obj.Put(2,3)
	obj.Put(4,1)
	param_1 =obj.Get(1)
	fmt.Println(param_1)

	//obj.Put(4,4)


	param_1 =obj.Get(2)
	fmt.Println(param_1)

	//param_1 =obj.Get(3)
	//fmt.Println(param_1)
	//
	//param_1 =obj.Get(4)
	//fmt.Println(param_1)
}
