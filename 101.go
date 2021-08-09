package main

func is_symmetric_tree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return is_symmetric_tree(p.Left, q.Right) && is_symmetric_tree(p.Right, q.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return is_symmetric_tree(root.Left, root.Right)
}
