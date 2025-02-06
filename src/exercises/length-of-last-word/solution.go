package solution

func lengthOfLastWord(s string) int {

	var count int
	var prevSpace bool

	for _, val := range s {
		if val == ' ' {
			prevSpace = true
			continue
		}

		if prevSpace {
			prevSpace = false
			count = 0
		}
		count++
	}

	return count
}
