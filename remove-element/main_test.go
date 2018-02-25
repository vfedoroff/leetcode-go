package removeelement

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		Nums []int
		Val  int
		Len  int
	}{
		{
			Nums: []int{3, 2, 2, 3},
			Val:  3,
			Len:  2,
		},
		{
			Nums: []int{1},
			Val:  1,
			Len:  0,
		},
		{
			Nums: []int{3, 3, 3},
			Val:  3,
			Len:  0,
		},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("nums=%v val=%d expected=%d", test.Nums, test.Val, test.Len), func(t *testing.T) {
			res := removeElement(test.Nums, test.Val)
			if res != test.Len {
				t.Errorf("wanted %d, but got %d", test.Len, res)
			}
		})
	}
}
