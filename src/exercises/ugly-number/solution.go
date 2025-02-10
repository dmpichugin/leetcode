package solution

func isUgly(n int) bool {
	if n < 1 {
		return false
	}

	if n == 1 {
		return true
	}

	prev := n
	for {
		if prev == 1 {
			return true
		}
		if prev%2 == 0 {
			prev /= 2
			continue
		}
		if prev%3 == 0 {
			prev /= 3
			continue
		}
		if prev%5 == 0 {
			prev /= 5
			continue
		}
		return false
	}
}
