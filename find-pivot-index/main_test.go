package leetcode

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	i := pivotIndex(nums)
	if i != 3 {
		t.Log(fmt.Sprintf("expected %d, but got %d", 3, i))
		t.FailNow()
	}

}
