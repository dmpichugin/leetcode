package solution

func isHappy(n int) bool {

	v := n
	m := make(map[int]int)
	for {

		digits := getSliceDigits(v)
		v = getSumSquares(digits)
		if v == 1 {
			return true
		}
		if v < 10 {
			m[v]++
			c := m[v]
			if c > 1 {
				break
			}
		}

	}

	return false
}

func getSumSquares(nums []int) int {

	value := 0
	for _, num := range nums {
		value += num * num
	}
	return value
}

func getSliceDigits(n int) []int {

	res := make([]int, 0, 9)

	tmp := n
	for tmp != 0 {

		a := tmp % 10
		res = append(res, a)
		tmp /= 10

	}

	return res
}
