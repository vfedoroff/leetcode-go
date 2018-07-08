package leecode

func addBinary(a string, b string) string {
	var ret []rune
	var sum rune
	ar := []rune(a)
	br := []rune(b)
	i := len(ar) - 1
	j := len(br) - 1
	for i >= 0 || j >= 0 || sum == 1 {
		if i >= 0 {
			sum = sum + ar[i] - '0'
		} else {
			sum = sum + 0
		}
		if j >= 0 {
			sum = sum + br[j] - '0'
		} else {
			sum = sum + 0
		}
		// If current digit sum is 1 or 3, add 1 to result
		ret = append([]rune{rune(sum%2 + '0')}, ret...)
		sum = sum / 2
		i--
		j--
	}
	return string(ret)
}
