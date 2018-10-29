package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func dfs(root *TreeNode, depth int, maxdepth *int) {
	if root == nil {
		return
	}
	if depth > *maxdepth {
		*maxdepth = depth
	}
	dfs(root.Left, depth+1, maxdepth)
	dfs(root.Right, depth+1, maxdepth)
}
func maxDepth(root *TreeNode) int {
	var maxdepth int = 0
	dfs(root, 1, &maxdepth)
	return maxdepth
}
