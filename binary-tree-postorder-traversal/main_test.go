package binarytreepostordertraversal

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	testCases := []struct {
		Input  *TreeNode
		Output []int
	}{
		{
			Input: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			Output: []int{3, 2, 1},
		},
		{
			Input:  nil,
			Output: []int{},
		},
		{
			Input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Output: []int{2, 1, 3},
		},
	}

	for _, test := range testCases {
		out := postorderTraversal(test.Input)
		if !reflect.DeepEqual(test.Output, out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
