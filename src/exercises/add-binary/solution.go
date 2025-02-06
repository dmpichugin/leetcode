package solution

func addBinary(a string, b string) string {

	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	s1 := []rune(a)
	s2 := []rune(b)

	result := make([]rune, max(len(s1), len(s2))+1)

	var overFlow bool
	i1 := len(s1) - 1
	i2 := len(s2) - 1

	var v1, v2 rune
	for i1 >= 0 || i2 >= 0 {
		v1, v2 = '0', '0'

		if i1 >= 0 {
			v1 = s1[i1]
		}
		if i2 >= 0 {
			v2 = s2[i2]
		}

		var sum int
		if v1 == '1' {
			sum += 1
		}
		if v2 == '1' {
			sum += 1
		}
		if overFlow {
			overFlow = false
			sum += 1
		}
		if sum > 1 {
			overFlow = true
			sum = sum % 2
		}

		var indexValue rune
		var writeIndex int

		writeIndex = max(i1, i2) + 1

		switch sum {
		case 0:
			indexValue = '0'
		case 1:
			indexValue = '1'
		}

		result[writeIndex] = indexValue

		if i1 >= 0 {
			i1--
		}
		if i2 >= 0 {
			i2--
		}
	}
	if overFlow {
		result[0] = '1'
	} else {
		result = result[1:]
	}

	return string(result)
}
