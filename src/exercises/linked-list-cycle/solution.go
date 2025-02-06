package solution

import (
	. "github.com/dmpichugin/leetcode/src/deps"
)

func hasCycle(head *ListNode) bool {

	p1 := head
	p2 := head

	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
		if p2 == nil {
			return false
		}

		if p1 == p2 {
			return true
		}
	}

	return false
}
