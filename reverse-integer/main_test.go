package reverseinteger

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		Input  int
		Output int
	}{
		{
			Input:  123,
			Output: 321,
		},
		{
			Input:  11,
			Output: 11,
		},
		{
			Input:  -1,
			Output: -1,
		},
		{
			Input:  -123,
			Output: -321,
		},
		{
			Input:  120,
			Output: 21,
		},
		{
			Input:  1534236469,
			Output: 0,
		},
		{
			Input:  0,
			Output: 0,
		},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("input=%d output=%d", test.Input, test.Output), func(t *testing.T) {
			res := reverse(test.Input)
			if res != test.Output {
				t.Errorf("wanted %d, but got %d", test.Output, res)
			}
		})
	}
}
