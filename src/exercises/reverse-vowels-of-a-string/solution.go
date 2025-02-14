package solution

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {

	s1 := []rune(s)

	l := 0
	r := len(s1) - 1

	for l < r {

		if !vowels[s1[l]] {
			l++
			continue
		}
		if !vowels[s1[r]] {
			r--
			continue
		}

		s1[l], s1[r] = s1[r], s1[l]
		l++
		r--
	}

	return string(s1)
}
