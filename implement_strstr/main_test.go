package leecode

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Haystack string
	Needle   string
	Output   int
}{
	{
		Haystack: "hello",
		Needle:   "ll",
		Output:   2,
	},
}

func TestStrStr(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s %s %d", tc.Haystack, tc.Needle, tc.Output), func(t *testing.T) {
			actual := strStr(tc.Haystack, tc.Needle)
			if actual != tc.Output {
				t.Errorf("expects: %d gets: %d", tc.Output, actual)
			}
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	for _, tc := range testCases {
		// run the Fib function b.N times
		b.Run(fmt.Sprintf("%s %s %d", tc.Haystack, tc.Needle, tc.Output), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				strStr(tc.Haystack, tc.Needle)
			}
		})
	}
}
