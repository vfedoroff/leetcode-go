package leecode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Output []int
	}{
		{
			Input:  [][]int{},
			Output: []int{},
		},
		{
			Input: [][]int{
				{1},
			},
			Output: []int{1},
		},
		{
			Input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Output: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v:%d", tc.Input, tc.Output), func(t *testing.T) {
			actual := findDiagonalOrder(tc.Input)
			if !reflect.DeepEqual(actual, tc.Output) {
				t.Errorf("expected %d, but get %d", tc.Output, actual)
			}
		})
	}
}
