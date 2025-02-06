package solution

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff1 := make(map[string]int)
	diff2 := make(map[string]int)

	sr1 := []rune(s1)
	sr2 := []rune(s2)

	diffCount := 0

	for i := 0; i < len(sr1); i++ {
		if diffCount > 2 {
			return false
		}
		if sr1[i] != sr2[i] {
			diff1[string(sr1[i])]++
			diff2[string(sr2[i])]++
			diffCount++
		}
	}

	for key, value := range diff1 {
		if value != diff2[key] || value != 1 || diffCount > 2 {
			return false
		}
	}

	return true
}
