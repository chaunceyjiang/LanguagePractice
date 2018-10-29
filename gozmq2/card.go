package main

/* import (
	"fmt"
)

type BSTreeNode struct {
	m_nValue int // value of node

	m_pLeft *BSTreeNode // left child of node

	m_pRight *BSTreeNode // right child of node

}

func main() {
	root:=new(BSTreeNode)
	root.m_nValue=10
	L:=[]int{6,4,8,14,12,16}
	for _,v:=range L{
		addNode(root, v)
	}
	printNode(root)
}
func addNode(root *BSTreeNode, value int) {
	node := new(BSTreeNode)
	node.m_nValue = value
	if value < root.m_nValue {
		if root.m_pLeft==nil{
			root.m_pLeft = node
		}else{
			addNode(root.m_pLeft, value)
		}
		
	} else {
		if root.m_pRight==nil{
			root.m_pRight = node
		}else{
			addNode(root.m_pRight, value)
		}
		
	}
}
func printNode(root *BSTreeNode){
	if root.m_pLeft!=nil{
		printNode(root.m_pLeft)
	}
	fmt.Println(root.m_nValue)
	if root.m_pRight!=nil{
		printNode(root.m_pRight)
	}

} */