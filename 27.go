package main

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		for j < len(nums) && nums[j] == val {
			j++
		}
		if j < len(nums) {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}
