package leetcode

func pivotIndex(nums []int) int {
	sum := 0
	for _, x := range nums {
		sum = sum + x
	}
	leftsum := 0
	for i, y := range nums {
		if leftsum == sum-leftsum-y {
			return i
		}
		leftsum = leftsum + y
	}
	return -1
}
