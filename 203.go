package main

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		for cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
