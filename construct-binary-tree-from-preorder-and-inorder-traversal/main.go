package constructbinarytree

// TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	val := preorder[preStart]
	node := &TreeNode{Val: val}
	//find parent element index from inorder
	k := 0
	for i, n := range inorder {
		if val == n {
			k = i
			break
		}
	}
	node.Left = construct(preorder, preStart+1, preStart+(k-inStart), inorder, inStart, k-1)
	node.Right = construct(preorder, preStart+(k-inStart)+1, preEnd, inorder, k+1, inEnd)
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preStart := 0
	preEnd := len(preorder) - 1
	inStart := 0
	inEnd := len(inorder) - 1
	return construct(preorder, preStart, preEnd, inorder, inStart, inEnd)
}
