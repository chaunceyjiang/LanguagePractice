package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findNode(root, p.Val, q.Val)
}
func findNode(root *TreeNode, p, q int) *TreeNode {
	if root.Val > p && root.Val > q {
		return findNode(root.Left, p, q)
	} else if root.Val < p && root.Val < q {
		return findNode(root.Right, p, q)
	}
	return root
}
func main() {

}
