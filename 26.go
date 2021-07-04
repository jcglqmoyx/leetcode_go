package main

func removeDuplicates(nums []int) int {
	if len(nums)== 0 {
		return 0
	}
	i, j := 0, 0
	for j < len(nums) {
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		if j < len(nums) {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
