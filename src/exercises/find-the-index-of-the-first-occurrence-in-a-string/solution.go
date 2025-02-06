package solution

func strStr(haystack string, needle string) int {

	s1 := []rune(haystack)
	s2 := []rune(needle)

	if len(s2) == 0 || len(s1) == 0 {
		return -1
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[0] {
			continue
		}

		var equal = true
		for j := 0; j < len(s2); j++ {
			index := i + j
			if index >= len(s1) {
				return -1
			}
			if s1[i+j] != s2[j] {
				equal = false
				break
			}
		}
		if equal {
			return i
		}
	}

	return -1
}
