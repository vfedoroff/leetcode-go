package pathsum

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	testCases := []struct {
		Input  *TreeNode
		Sum    int
		Output bool
	}{
		{
			Input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			Sum:    22,
			Output: true,
		},
	}

	for _, test := range testCases {
		out := hasPathSum(test.Input, test.Sum)
		if !(test.Output == out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
