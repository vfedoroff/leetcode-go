package binarytreepreordertraversal

// TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack data structure that implements LIFO
type Stack []*TreeNode

// Push adds an item onto the stack.
func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

// Pop removes the most-recently-pushed item from the stack.
func (s *Stack) Pop() *TreeNode {
	arr := *s
	last := len(arr) - 1
	top := arr[last]
	if len(arr) > 0 {
		arr = arr[0:last]
	} else {
		arr = []*TreeNode{}
	}
	*s = arr
	return top
}

// Nice explanation: https://www.geeksforgeeks.org/iterative-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := Stack{root}
	for len(stack) > 0 {
		node := stack.Pop()
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
	return ret
}

func preorderTraversalRecursion(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	var innerFunc func(root *TreeNode)
	innerFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		// now visiting root
		ret = append(ret, root.Val)
		// first recur on left subtree
		innerFunc(root.Left)
		// then recur on right subtree
		innerFunc(root.Right)
	}
	innerFunc(root)
	return ret
}
