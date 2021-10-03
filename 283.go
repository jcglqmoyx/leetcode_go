package main

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			i++
		}
	}
}
