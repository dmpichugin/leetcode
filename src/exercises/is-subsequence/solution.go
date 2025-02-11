package solution

func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	s1 := []rune(s)
	s2 := []rune(t)

	index := 0

	for i := 0; i < len(s2); i++ {
		if index < len(s1) && s2[i] == s1[index] {
			index++
		}
	}
	return index == len(s1)
}
