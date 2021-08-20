package main

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func dfs(node *TreeNode, flag *bool) int {
	if !(*flag) {
		return 0
	}
	if node == nil {
		return 0
	}
	l, r := dfs(node.Left, flag), dfs(node.Right, flag)
	if abs(l-r) > 1 {
		*flag = false
	}
	return max(l, r) + 1
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var flag bool = true
	dfs(root, &flag)
	return flag
}
