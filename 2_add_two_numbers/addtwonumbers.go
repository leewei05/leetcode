package addtwonumbers

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// O(max(m,n)) m for l1 length, n for l2 length
// O(max(m,n)) +1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	first := res
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		res.Next = &ListNode{sum % 10, nil}
		res = res.Next
	}

	if carry > 0 {
		res.Next = &ListNode{carry, nil}
	}

	return first.Next
}
