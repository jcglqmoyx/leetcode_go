package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				head.Next = l1
				l1 = l1.Next
			} else {
				head.Next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	return res.Next
}
