package main

func firstMissingPositive(nums []int) int {
	var abs = func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	for _, v := range nums {
		v = abs(v)
		if v <= n {
			nums[v-1] = -abs(nums[v-1])
		}
	}
	for i := range nums {
		if nums[i] >= 0 {
			return i + 1
		}
	}
	return n + 1
}
