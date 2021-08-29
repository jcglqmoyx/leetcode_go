package main

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{-1, head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	right := slow.Next
	slow.Next = nil
	right = reverse(right)
	for head != nil && right != nil {
		if head.Val != right.Val {
			return false
		}
		head = head.Next
		right = right.Next
	}
	return true
}
