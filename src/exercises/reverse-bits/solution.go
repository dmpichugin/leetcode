package solution

func reverseBits(num uint32) uint32 {

	result := uint32(0)
	for i := 0; i <= 31; i++ {
		result = result << 1
		result = result | ((num >> (i)) & 1)
	}

	return result
}
