package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val >= p.Val && root.Val <= q.Val || root.Val <= p.Val && root.Val >= q.Val {
		return root
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return lowestCommonAncestor(root.Left, p, q)
}
