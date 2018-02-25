package atoi

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		Input  string
		Output int
	}{
		{
			Input:  "",
			Output: 0,
		},
		{
			Input:  "-",
			Output: 0,
		},
		{
			Input:  "+",
			Output: 0,
		},
		{
			Input:  "0",
			Output: 0,
		},
		{
			Input:  "123",
			Output: 123,
		},
		{
			Input:  "-1",
			Output: -1,
		},
		{
			Input:  "1",
			Output: 1,
		},
		{
			Input:  "2147483648",
			Output: 2147483647,
		},
		{
			Input:  strconv.Itoa(math.MinInt32),
			Output: -2147483648,
		},
		{
			Input:  "    010",
			Output: 10,
		},
		{
			Input:  "1a",
			Output: 1,
		},
		{
			Input:  "  -0012a42",
			Output: -12,
		},
		{
			Input:  "  -a1",
			Output: 0,
		},
		{
			Input:  "9223372036854775809",
			Output: 2147483647,
		},
		{
			Input:  "             ",
			Output: 0,
		},
		{
			Input:  "   -1123u3761867",
			Output: -1123,
		},
		{
			Input:  "      -11919730356x",
			Output: -2147483648,
		},
		{
			Input:  "  -1010023630o4",
			Output: -1010023630,
		},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("input=%s output=%d", test.Input, test.Output), func(t *testing.T) {
			res := myAtoi(test.Input)
			if res != test.Output {
				t.Errorf("wanted %d, but got %d", test.Output, res)
			}
		})
	}
}
