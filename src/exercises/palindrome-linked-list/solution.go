package solution

import . "github.com/dmpichugin/leetcode/src/deps"

func isPalindrome(head *ListNode) bool {

	var count int
	last := head

	for last != nil {
		count++
		last = last.Next
	}

	var skipElement int = -1

	if count%2 == 1 {
		skipElement = count/2 + 1
	}

	stack := make([]int, 0, count/2)
	middle := count / 2
	index := 1

	cursor := head

	for cursor != nil {
		if index == skipElement {
			cursor = cursor.Next
			index++
			continue
		}

		if index <= middle {
			stack = append(stack, cursor.Val)
		} else {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if val != cursor.Val {
				return false
			}
		}
		index++
		cursor = cursor.Next
	}

	return true
}
