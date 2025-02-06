package solution

import (
	"unicode"
)

func isPalindrome(s string) bool {

	s1 := []rune(s)

	l := 0
	r := len(s1) - 1

	for l <= r {
		if !isLetter(s1[l]) {
			l++
			continue
		}
		if !isLetter(s1[r]) {
			r--
			continue
		}

		if !isEqual(s1[l], s1[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isLetter(c1 rune) bool {
	return unicode.IsLetter(c1) || unicode.IsDigit(c1)
}

func isEqual(c1, c2 rune) bool {
	return unicode.ToLower(c1) == unicode.ToLower(c2)
}
