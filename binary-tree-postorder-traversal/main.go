package binarytreepostordertraversal

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

// Peek a top element
func (s *Stack) Peek() *TreeNode {
	arr := *s
	if len(arr) == 0 {
		return nil
	}
	last := len(arr) - 1
	top := arr[last]
	return top
}

// Nice explanation: https://www.geeksforgeeks.org/iterative-postorder-traversal-using-stack/
func postorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	stack := Stack{}
	for {
		// Push root's right child and then root to stack
		for root != nil {
			if root.Right != nil {
				stack.Push(root.Right)
			}
			stack.Push(root)
			// Set root as root's left child
			root = root.Left
		}
		// Pop an item from stack and set it as root
		root = stack.Pop()
		// If the popped item has a right child and the
		// right child is not processed yet, then make sure
		// right child is processed before root
		if root.Right != nil && stack.Peek() == root.Right {
			stack.Pop()       //Remove right child from stack
			stack.Push(root)  //Push root back to stack
			root = root.Right // change root so that the right childis processed next
		} else { // Else print root's data and set root as None
			ret = append(ret, root.Val)
			root = nil
		}
		if len(stack) <= 0 {
			break
		}
	}
	return ret
}
