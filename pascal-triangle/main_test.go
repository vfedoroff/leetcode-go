package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	actual := generate(5)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expects: %v, but gets: %v", expected, actual)
	}
}
