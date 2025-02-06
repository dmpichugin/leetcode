package solution

func plusOne(digits []int) []int {

	var overFlow bool

	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] += 1
		}

		if overFlow {
			digits[i] += 1
			overFlow = false
		}

		if digits[i] > 9 {
			overFlow = true
			digits[i] %= 10
		}
	}

	if overFlow {
		res := make([]int, 0, len(digits)+1)
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}

	return digits
}
