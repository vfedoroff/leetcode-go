package leecode

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func dominantIndex(nums []int) int {
	maxVal := minInt
	secondMax := minInt
	var maxIndex int
	for i, x := range nums {
		if x > maxVal {
			secondMax = maxVal
			maxVal = x
			maxIndex = i
			continue
		}
		if x > secondMax {
			secondMax = x
		}
	}
	if maxVal-secondMax >= secondMax {
		return maxIndex
	}
	return -1
}
