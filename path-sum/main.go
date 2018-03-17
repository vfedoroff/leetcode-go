package pathsum

// TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && (root.Left == nil && root.Right == nil) {
		return true
	}
	var ret = hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	return ret
}
