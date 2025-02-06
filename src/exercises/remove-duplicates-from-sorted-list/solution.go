package solution

import (
	. "github.com/dmpichugin/leetcode/src/deps"
)

func deleteDuplicates(head *ListNode) *ListNode {

	cursor := head
	if cursor == nil || cursor.Next == nil {
		return head
	}

	for cursor.Next != nil {

		if cursor.Val == cursor.Next.Val {
			cursor.Next = cursor.Next.Next
		} else {
			cursor = cursor.Next
		}
	}

	return head
}
