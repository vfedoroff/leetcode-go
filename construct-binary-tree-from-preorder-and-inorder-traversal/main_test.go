package constructbinarytree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		Preorder []int
		Inorder  []int
		Output   *TreeNode
	}{
		{
			Preorder: []int{3, 9, 20, 15, 7},
			Inorder:  []int{9, 3, 15, 20, 7},
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
		out := buildTree(test.Preorder, test.Inorder)
		if !reflect.DeepEqual(test.Output, out) {
			t.Errorf("expected to get %#v, but got %#v", test.Output, out)
		}
	}
}
