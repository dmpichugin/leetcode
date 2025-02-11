package solution

import (
	. "github.com/dmpichugin/leetcode/src/deps"
)

func removeElements(head *ListNode, val int) *ListNode {

	var resHead, prevElement *ListNode

	cursor := head
	for cursor != nil {
		if cursor.Val != val {
			if resHead == nil {
				resHead = cursor
			}

			prevElement = cursor
			cursor = cursor.Next
			continue
		} else {
			if prevElement != nil {
				prevElement.Next = cursor.Next
			}
			cursor = cursor.Next
		}
	}

	return resHead
}
