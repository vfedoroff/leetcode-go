package leecode

const base = 16777619

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func strStr(haystack string, needle string) int {
	n := len(haystack) // string len
	m := len(needle)   // pattern len
	switch {
	case m == 0:
		return m
	case n == m:
		if haystack == needle {
			return 0
		}
		return -1
	case m > n:
		return -1
	}
	pattern := hash(needle)

	for i := 0; i < (n-m)+1; i++ {
		substr := haystack[i : i+m]
		hs := hash(substr)
		if hs == pattern {
			return i
		}
	}
	return -1
}
