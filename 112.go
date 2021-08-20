package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}
