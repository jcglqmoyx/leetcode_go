package main

import "strconv"

var res []string

func dfs(root *TreeNode, cur string) {
	if root == nil {
		return
	}
	tmp := cur
	tmp += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, tmp)
	} else {
		tmp += "->"
		dfs(root.Left, tmp)
		dfs(root.Right, tmp)
	}
}
func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	dfs(root, "")
	return res
}
