package main

func numberOfArithmeticSlices(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i+1]*2 == nums[i]+nums[i+2] {
			j := i + 1
			for j < len(nums)-2 {
				if nums[j+1]*2 == nums[j]+nums[j+2] {
					j++
				} else {
					break
				}
			}
			length := j - i + 2
			cnt += (length - 2) * (length - 1) / 2
			i = j
		}
	}
	return cnt
}
