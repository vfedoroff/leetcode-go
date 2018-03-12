package maximumdepthofbunarytree

// TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	ret := max(left, right) + 1
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
