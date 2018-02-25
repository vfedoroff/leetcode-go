package removedublicatesfromsortedarray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	res := removeDuplicates([]int{1, 1, 2})
	if res != 2 {
		t.Fail()
	}
}
