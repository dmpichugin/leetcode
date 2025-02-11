package solution

func countBits(n int) []int {
	var results []int
	for i := 0; i <= n; i++ {
		bitsCount := getBits(i)
		results = append(results, bitsCount)
	}

	return results
}

func getBits(num int) int {

	var count int
	for i := 0; i < 32; i++ {
		if num>>i&1 == 1 {
			count++
		}
	}

	return count
}
