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

// Nice explanation: https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion/
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := Stack{}
	current := root
	for {
		if current != nil {
			stack.Push(current)
			current = current.Left
			continue
		}
		if len(stack) > 0 {
			current = stack.Pop()
			ret = append(ret, current.Val)
			current = current.Right
			continue
		}
		break
	}
	return ret
}

func inorderTraversalRecursion(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	var innerFunc func(root *TreeNode)
	innerFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		// first recur on left subtree
		innerFunc(root.Left)
		// now visiting root
		ret = append(ret, root.Val)
		// then recur on right subtree
		innerFunc(root.Right)
	}
	innerFunc(root)
	return ret
}
