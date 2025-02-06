package solution

func isPowerOfTwo(n int) bool {
	if n == 1 || n == 2 {
		return true
	}
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
