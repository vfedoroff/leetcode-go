package binarytreepostordertraversal

import (
	"reflect"
	"testing"
)

func TestLevelorderTraversal(t *testing.T) {
	testCases := []struct {
		Input  *TreeNode
		Output [][]int
	}{
		{
			Input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Output: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			Input:  nil,
			Output: [][]int{},
		},
		{
			Input: &TreeNode{
				Val: 9,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: -7,
							Left: &TreeNode{
								Val: -1,
							},
						},
					},
				},
			},
			Output: [][]int{
				{9},
				{2},
				{0},
				{-7},
				{-1},
			},
		},
	}

	for _, test := range testCases {
		out := levelOrder(test.Input)
		if !reflect.DeepEqual(test.Output, out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
