package permutations

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		Input  []int
		Output [][]int
	}{
		{
			Input: []int{1, 2, 3},
			Output: [][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
		},
		{
			Input: []int{1},
			Output: [][]int{
				[]int{1},
			},
		},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("input=%v output=%v", test.Input, test.Output), func(t *testing.T) {
			output := permute(test.Input)
			if len(output) != len(test.Output) {
				t.Errorf("wanted len %d, but got %d", len(test.Output), len(output))
			}
			sort.Slice(output, func(i, j int) bool {
				a, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output[i])), ""), "[]"))
				b, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output[j])), ""), "[]"))
				return a < b
			})
			if !reflect.DeepEqual(test.Output, output) {
				t.Errorf("wanted %v, but got %v", test.Output, output)
			}
		})
	}
}
