package twosum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		complement := target - v
		c, ok := m[complement]
		if ok {
			return []int{c, i}
		}
		m[v] = i
	}
	panic("no solution")
}
