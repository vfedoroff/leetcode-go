package leecode

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if m == 1 && n == 1 {
		return matrix[0]
	}
	ret := []int{}
	var row, col int
	for i := 0; i < m*n; i++ {
		ret = append(ret, matrix[row][col])
		if (row+col)%2 == 0 { // moving up
			if col == n-1 {
				row++
				continue
			}
			if row == 0 {
				col++
				continue
			}
			row--
			col++
		} else { // moving down
			if row == m-1 {
				col++
				continue
			}
			if col == 0 {
				row++
				continue
			}
			row++
			col--
		}
	}
	return ret
}
