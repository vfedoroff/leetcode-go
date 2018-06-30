package leecode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDominantIndex(t *testing.T) {
	testCases := []struct {
		Input  []int
		Output []int
	}{
		{
			Input:  []int{9},
			Output: []int{1, 0},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v:%d", tc.Input, tc.Output), func(t *testing.T) {
			actual := plusOne(tc.Input)
			if !reflect.DeepEqual(actual, tc.Output) {
				t.Errorf("expected %d, but get %d", tc.Output, actual)
			}
		})
	}
}
