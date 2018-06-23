package constructbinarytree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		Inorder   []int
		Postorder []int
		Output    *TreeNode
	}{
		{
			Inorder:   []int{9, 3, 15, 20, 7},
			Postorder: []int{9, 15, 7, 20, 3},
			Output: &TreeNode{
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
		},
	}

	for _, test := range testCases {
		out := buildTree(test.Inorder, test.Postorder)
		if !reflect.DeepEqual(test.Output, out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
