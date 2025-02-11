package deps

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntList(nums ...int) *ListNode {

	return GetList(nums...)
}

func GetList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var head, lastElement *ListNode

	for _, num := range nums {
		if head == nil || lastElement == nil {
			head = &ListNode{
				Val: num,
			}
			lastElement = head
		} else {
			lastElement.Next = &ListNode{
				Val: num,
			}
			lastElement = lastElement.Next
		}
	}
	return head
}

func (i *ListNode) String() string {

	sb := strings.Builder{}
	sb.WriteString("[")
	head := i
	for head != nil {
		sb.WriteString(fmt.Sprint(head.Val))
		if head.Next != nil {
			sb.WriteString(",")
		}
		head = head.Next
	}
	sb.WriteString("]")
	return sb.String()
}

func (i *ListNode) MakeCycle() *ListNode {
	if i == nil {
		return nil
	}

	cursor := i
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = i
	return i
}

func (i *ListNode) MakeCycleInMiddle() *ListNode {
	if i == nil {
		return nil
	}

	var count = 1
	cursor := i
	for cursor.Next != nil {
		cursor = cursor.Next
		count++
	}

	lastElement := cursor

	middleCursor := i

	count /= 2
	for count >= 0 {
		middleCursor = middleCursor.Next
		count--
	}

	if lastElement.Next == nil {
		lastElement.Next = middleCursor
	}

	return i
}

func (i *ListNode) Equals(b *ListNode) bool {

	l1 := i
	l2 := b

	if l1 == nil && l2 != nil {
		return false
	}
	if l1 != nil && l2 == nil {
		return false
	}
	if l1 == nil && l2 == nil {
		return true
	}

	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			return false
		}
		if l1 != nil && l2 == nil {
			return false
		}

		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}
