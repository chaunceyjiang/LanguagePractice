package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var flag bool
	dfs(root, &k, &res, &flag)
	return res
}

func dfs(root *TreeNode, k *int, res *int, flag *bool) {
	if root == nil{
		return
	}

	dfs(root.Left,k,res,flag)
		*k = (*k) -1
	if *k == 0 && *flag == false{
		*flag = true
		*res = root.Val
	}

	dfs(root.Right,k,res,flag)
}

func main() {
	kthSmallest()
}
