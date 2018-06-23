package constructbinarytree

// TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if (len(inorder) != len(postorder)) || (len(inorder) == 0 || len(postorder) == 0) {
		return nil
	}
	// Root is the last node in postorder[]
	ret := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	if len(inorder) == 1 {
		return ret
	}
	idx := indexOf(ret.Val, inorder)
	// Everything on left of “idx” is in left subtree
	ret.Left = buildTree(inorder[:idx], postorder[:idx])
	// Everything on right of “idx” is in right subtree
	ret.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return ret
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}
