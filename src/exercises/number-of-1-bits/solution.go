package solution

func hammingWeight(n int) int {

	var count int
	for i := 0; i < 32; i++ {
		if (n>>(31-i))&1 == 1 {
			count++
		}
	}

	return count
}
