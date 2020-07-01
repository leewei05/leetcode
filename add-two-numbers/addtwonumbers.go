package addtwonumbers

// ListNode is a struct of a linked-list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{}

	for l1 != nil && l2 != nil {
		res.Val = l1.Val + l2.Val
		res.Next = &ListNode{}
		l1 = l1.Next
		l2 = l2.Next
	}

	return &res
}
