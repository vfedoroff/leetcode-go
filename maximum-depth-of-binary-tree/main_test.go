package maximumdepthofbunarytree

import (
	"reflect"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		Input  *TreeNode
		Output int
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
			Output: 3,
		},
		{
			Input:  nil,
			Output: 0,
		},
	}

	for _, test := range testCases {
		out := maxDepth(test.Input)
		if !reflect.DeepEqual(test.Output, out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
