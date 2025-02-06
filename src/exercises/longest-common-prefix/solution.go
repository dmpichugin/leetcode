package solution

import "strings"

func longestCommonPrefix(strs []string) string {

	parsed := make([][]rune, 0, len(strs))
	maxLen := 0
	for _, str := range strs {
		parsedValue := []rune(str)
		parsed = append(parsed, parsedValue)
		maxLen = max(maxLen, len(parsedValue))
	}

	sb := strings.Builder{}

	var r rune
LOOP:
	for j := 0; j < maxLen; j++ {
		r = 0

		for i := 0; i < len(parsed); i++ {
			if j >= len(parsed[i]) {
				break LOOP
			}

			if i == 0 {
				r = parsed[i][j]
			}

			if parsed[i][j] != r {
				break LOOP
			}
		}
		sb.WriteRune(r)
	}

	return sb.String()
}
