package solution

func addDigits(num int) int {
	v := num
	for v > 9 {
		v = sumSlice(getDigits(v))
	}
	return v
}

func getDigits(num int) []int {

	v := num
	var digits []int
	var tmp int
	for v != 0 {
		tmp = v % 10
		digits = append(digits, tmp)
		v /= 10
	}

	return digits

}

func sumSlice(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
