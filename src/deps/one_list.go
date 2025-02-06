package deps

import (
	"fmt"
	"strings"
)

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func GetIntList(nums ...int) *ListNode[int] {

	return GetList(nums...)
}

func GetList[T any](nums ...T) *ListNode[T] {
	if len(nums) == 0 {
		return nil
	}

	var head, lastElement *ListNode[T]

	for _, num := range nums {
		if head == nil || lastElement == nil {
			head = &ListNode[T]{
				Val: num,
			}
			lastElement = head
		} else {
			lastElement.Next = &ListNode[T]{
				Val: num,
			}
			lastElement = lastElement.Next
		}
	}
	return head
}

func (i *ListNode[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	head := i
	sb.WriteString(fmt.Sprint(head.Val))
	for head != nil {
		if head.Next != nil {
			sb.WriteString(",")
		}
		head = head.Next
	}
	sb.WriteString("]")
	return sb.String()
}
