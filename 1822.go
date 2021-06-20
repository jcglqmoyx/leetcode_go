package main

func arraySign(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			cnt++
		} else if nums[i] == 0 {
			return 0
		}
	}
	if cnt&1 == 1 {
		return -1
	}
	return 1
}
