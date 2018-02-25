package atoi

import (
	"strings"
)

const intMax = 2147483647
const intMin = -2147483648

func myAtoi(str string) int {
	prefix := 1
	str = strings.TrimSpace(str)
	l := len(str)
	if l == 0 {
		return 0
	}
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			prefix = -1
		}
		str = str[1:]
		l = len(str)
		if l < 1 {
			return 0
		}
	}
	num := 0
	i := 0
	for {
		c := byte(str[i])
		c = c - '0'
		if c > 9 {
			break
		}
		num = num*10 + int(c)
		i++
		if i > 10 && i < l {
			if prefix > 0 {
				return intMax
			}
			return intMin
		}
		if i >= l {
			break
		}
	}
	res := prefix * num
	if res >= intMax {
		return intMax
	}
	if res <= intMin {
		return intMin
	}
	return prefix * num
}
