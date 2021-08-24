package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	flagA, flagB := false, false
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
		if headA == nil {
			if flagA {
				return nil
			}
			flagA = true
			headA = b
		}
		if headB == nil {
			if flagB {
				return nil
			}
			flagB = true
			headB = a
		}
	}
	return headA
}
