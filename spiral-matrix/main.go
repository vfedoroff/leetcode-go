package leecode

func spiralOrder(matrix [][]int) []int {
	ret := []int{}
	if len(matrix) == 0 {
		return ret
	}
	m := len(matrix)
	n := len(matrix[0])
	x := 0
	y := 0
	for m > 0 && n > 0 {
		//if one row/column left, no circle can be formed
		if m == 1 {
			for i := 0; i < n; i++ {
				ret = append(ret, matrix[x][y])
				y++
			}
			break
		} else if n == 1 {
			for i := 0; i < m; i++ {
				ret = append(ret, matrix[x][y])
				x++
			}
			break
		}

		//below, process a circle

		//top - move right
		for i := 0; i < n-1; i++ {
			ret = append(ret, matrix[x][y])
			y++
		}

		//right - move down
		for i := 0; i < m-1; i++ {
			ret = append(ret, matrix[x][y])
			x++
		}

		//bottom - move left
		for i := 0; i < n-1; i++ {
			ret = append(ret, matrix[x][y])
			y--
		}

		//left - move up
		for i := 0; i < m-1; i++ {
			ret = append(ret, matrix[x][y])
			x--
		}

		x++
		y++
		m = m - 2
		n = n - 2
	}
	return ret
}
