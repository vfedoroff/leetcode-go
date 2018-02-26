package permutations

// Good to read:
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permute(nums []int) [][]int {
	output := [][]int{}
	n := len(nums)
	permutation := make([]int, n)
	copy(permutation, nums)
	output = append(output, permutation)
	i := 0
	c := make([]int, n)
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				nums[0], nums[i] = nums[i], nums[0]
			} else {
				nums[c[i]], nums[i] = nums[i], nums[c[i]]
			}
			// We need to create a copy of the original array,
			// because slice keeps a references on values
			// each time when we do a swap we change all resulting permutaions
			permutation := make([]int, n)
			copy(permutation, nums)
			output = append(output, permutation)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return output
}
