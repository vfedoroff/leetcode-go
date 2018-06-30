package leecode

func plusOne(digits []int) []int {
	plusOne := true
	// start from the last element
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if plusOne {
		digits = append([]int{1}, digits...)
	}
	return digits
}
