package main

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := ListNode{0, head}
	cur := head
	for cur != nil {
		node := cur
		for node != nil && node.Val == cur.Val {
			node = node.Next
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}
