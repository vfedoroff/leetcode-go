package binarytreepostordertraversal

import (
	"fmt"
)

// TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("%d", t.Val)
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	var innerFunc func(level []*TreeNode)
	innerFunc = func(level []*TreeNode) {
		currentLevel := []int{}
		nextLevel := []*TreeNode{}
		if len(level) > 0 {
			for _, node := range level {
				if node != nil {
					currentLevel = append(currentLevel, node.Val)
					if node.Left != nil {
						nextLevel = append(nextLevel, node.Left)
					}
					if node.Right != nil {
						nextLevel = append(nextLevel, node.Right)
					}
				}
			}
			if len(currentLevel) > 0 {
				ret = append(ret, currentLevel)
			}
			innerFunc(nextLevel)
		}
	}
	innerFunc([]*TreeNode{root})
	return ret
}
