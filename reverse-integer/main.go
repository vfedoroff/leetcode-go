package reverseinteger

import (
	"math"
)

// Psevdo code:
// a. Extract off the rightmost digit of your input number. (1234 % 10) = 4
// b. Take that digit (4) and add it into a new reversedNum.
// c. Multiply reversedNum by 10 (4 * 10) = 40, this exposes a zero to the right of your (4).
// d. Divide the input by 10, (removing the rightmost digit). (1234 / 10) = 123
// e. Repeat at step a with 123
// Source: https://stackoverflow.com/questions/3806126/java-reverse-an-int-value-without-using-array
func reverse(x int) int {
	var n int
	negative := x < 0
	if negative {
		x = -x
	}
	for x != 0 {
		remainder := x % 10
		n = n*10 + remainder
		x = x / 10
	}
	if n > math.MaxInt32 {
		return 0
	}
	if negative {
		return -n
	}
	return n
}
