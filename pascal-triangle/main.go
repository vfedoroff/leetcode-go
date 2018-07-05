package leetcode

func generate(numRows int) [][]int {
	temp := 1
	var ret = [][]int{}
	for i := 0; i < numRows; i++ {
		var row = []int{}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			row = append(row, temp)
		}
		ret = append(ret, row)
	}
	return ret
}
