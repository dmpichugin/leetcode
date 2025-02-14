package solution

import (
	"sort"
)

func findTheDifference(s string, t string) byte {

	s1 := []byte(s)
	s2 := []byte(t)

	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	index := 0

	for i := 0; i < len(s2); i++ {
		if index >= len(s1) {
			return s2[i]
		}

		if s1[index] != s2[i] {
			return s2[i]
		}
		index++
	}

	return 0
}
