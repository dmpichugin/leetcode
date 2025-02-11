package solution

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	const mask = 1431655765 // in binary 01010101010101010101010101010101

	return ((n-1)&n) == 0 && n&mask != 0
}
