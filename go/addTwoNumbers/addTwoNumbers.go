package addTwoNumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode
	var sum = 0
	var carry = 0

	for l1 != nil && l2 != nil {
		sum = carry + l1.Val + l2.Val
		carry = sum / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{sum, nil}
			current = head
		} else {
			current.Next = &ListNode{sum, nil}
			current = current.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum = carry + l1.Val
		carry = sum / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{sum, nil}
			current = head
		} else {
			current.Next = &ListNode{sum, nil}
			current = current.Next
		}
		l1 = l1.Next
	}

	for l2 != nil {
		sum = carry + l2.Val
		carry = sum / 10
		sum = sum % 10
		if head == nil {
			head = &ListNode{sum, nil}
			current = head
		} else {
			current.Next = &ListNode{sum, nil}
			current = current.Next
		}
		l2 = l2.Next
	}

	if carry != 0 {
		current.Next = &ListNode{carry, nil}
	}

	return head
}
