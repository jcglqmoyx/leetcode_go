package main

//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
//func dfs(node *TreeNode, depth int, minD *int) {
//	if node == nil {
//		return
//	}
//	if node.Left == nil && node.Right == nil {
//		*minD = min(*minD, depth)
//	}
//	dfs(node.Left, depth+1, minD)
//	dfs(node.Right, depth+1, minD)
//}
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var minD = 100000
//	dfs(root, 1, &minD)
//	return minD
//}
