package solution

func isIsomorphic(s string, t string) bool {

	s1 := []rune(s)
	s2 := []rune(t)

	if len(s1) != len(s2) {
		return false
	}

	reverseM1 := make(map[rune]rune)
	reverseM2 := make(map[rune]rune)

	for i := 0; i < len(s1); i++ {

		origin1 := s1[i]
		origin2 := s2[i]

		reversed1, ok1 := reverseM1[origin1]
		reversed2, ok2 := reverseM2[origin2]

		if !ok1 && !ok2 {
			reverseM1[origin1] = origin2
			reverseM2[origin2] = origin1
			continue
		}

		if ok1 && ok2 {
			if origin1 != reversed2 {
				return false
			}
			if origin2 != reversed1 {
				return false
			}
			continue
		}
		return false
	}

	return true
}
