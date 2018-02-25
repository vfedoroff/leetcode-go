package removedublicatesfromsortedarray

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	j := 0
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}
	nums[j] = nums[n-1]
	j++
	return j
}
