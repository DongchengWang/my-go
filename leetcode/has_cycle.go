package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	l1 := head
	l2 := head.Next

	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return true
		}

		l1 = l1.Next
		l2 = l2.Next
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return false
}
