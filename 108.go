package main

func build(nums []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	root := &TreeNode{}
	root.Val = nums[mid]
	root.Left = build(nums, l, mid-1)
	root.Right = build(nums, mid+1, r)
	return root
}
func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}
