package main

func maxSubArray(nums []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res, sum := -1000000, 0
	for _, num := range nums {
		sum = max(sum+num, num)
		res = max(res, sum)
	}
	return res
}
