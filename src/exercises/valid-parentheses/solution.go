package solution

func isValid(s string) bool {

	var stack []rune

	for _, r := range s {

		switch r {
		case '{', '[', '(':
			stack = append(stack, r)
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}

			lastElement := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch r {
			case '}':
				if lastElement != '{' {
					return false
				}
			case ']':
				if lastElement != '[' {
					return false
				}
			case ')':
				if lastElement != '(' {
					return false
				}
			}

		default:
			return false
		}
	}
	return len(stack) == 0
}
