package leecode

import "testing"

func TestAddBinary(t *testing.T) {
	testcases := []struct {
		A        string
		B        string
		Expected string
	}{
		{
			A:        "11",
			B:        "1",
			Expected: "100",
		},
		{
			A:        "1010",
			B:        "1011",
			Expected: "10101",
		},
	}

	for _, tc := range testcases {
		actual := addBinary(tc.A, tc.B)
		if actual != tc.Expected {
			t.Errorf("not equal. Expected: %s, when get %s", tc.Expected, actual)
		}
	}
}
