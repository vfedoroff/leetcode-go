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
			Input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			Output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v:%d", tc.Input, tc.Output), func(t *testing.T) {
			actual := spiralOrder(tc.Input)
			if !reflect.DeepEqual(actual, tc.Output) {
				t.Errorf("expected %d, but get %d", tc.Output, actual)
			}
		})
	}
}
