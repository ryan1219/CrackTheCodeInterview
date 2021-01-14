package addtwonumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	result := make([]int, 0)
	var head *ListNode
	current := head
	carry := 0

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	for len(s1) != 0 && len(s2) != 0 {
		v1 := s1[len(s1)-1]
		v2 := s2[len(s2)-1]
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
		sum := carry + v1 + v2
		carry = sum / 10
		sum = sum % 10

		result = append(result, sum)
	}

	for len(s1) != 0 {
		v1 := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		sum := carry + v1
		carry = sum / 10
		sum = sum % 10

		result = append(result, sum)
	}

	for len(s2) != 0 {
		v2 := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		sum := carry + v2
		carry = sum / 10
		sum = sum % 10

		result = append(result, sum)
	}

	if carry == 1 {
		result = append(result, carry)
	}

	for len(result) != 0 {
		if head == nil {
			head = &ListNode{result[len(result)-1], nil}
			current = head
		} else {
			current.Next = &ListNode{result[len(result)-1], nil}
			current = current.Next
		}
		result = result[:len(result)-1]
	}
	return head
}
