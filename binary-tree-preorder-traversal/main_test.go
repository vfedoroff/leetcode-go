package binarytreepreordertraversal

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
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
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			Output: []int{1, 2, 3},
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
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Output: []int{3, 1, 2},
		},
	}

	for _, test := range testCases {
		out := preorderTraversal(test.Input)
		if !reflect.DeepEqual(test.Output, out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
