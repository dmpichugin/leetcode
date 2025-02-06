package solution

import (
	. "github.com/dmpichugin/leetcode/src/deps"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var head *ListNode
	var last *ListNode

	l1 := list1
	l2 := list2

	for l1 != nil || l2 != nil {
		var v int
		if l1 != nil && l2 == nil {
			v = l1.Val
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			v = l2.Val
			l2 = l2.Next
		} else {
			if l1.Val < l2.Val {
				v = l1.Val
				l1 = l1.Next
			} else {
				v = l2.Val
				l2 = l2.Next
			}
		}

		newVal := &ListNode{
			Val: v,
		}
		if head == nil {
			head = newVal
			last = head
		} else {
			last.Next = newVal
			last = newVal
		}
	}

	return head
}
