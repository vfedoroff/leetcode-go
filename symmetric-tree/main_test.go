package symmetrictree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		Input  *TreeNode
		Output bool
	}{
		{
			Input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			Output: true,
		},
	}

	for _, test := range testCases {
		out := isSymmetric(test.Input)
		if !(test.Output == out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
