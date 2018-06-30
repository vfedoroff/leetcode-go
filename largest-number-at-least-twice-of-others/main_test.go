package leecode

import (
	"fmt"
	"testing"
)

func TestDominantIndex(t *testing.T) {
	testCases := []struct {
		Input  []int
		Output int
	}{
		{
			Input:  []int{3, 6, 1, 0},
			Output: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v:%d", tc.Input, tc.Output), func(t *testing.T) {
			actual := dominantIndex(tc.Input)
			if actual != tc.Output {
				t.Errorf("expected %d, but get %d", tc.Output, actual)
			}
		})
	}
}
