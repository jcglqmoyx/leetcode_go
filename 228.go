package main

import "strconv"

func summaryRanges(nums []int) []string {
	var res []string
	for i := 0; i < len(nums); i++ {
		res = append(res, strconv.Itoa(nums[i]))
		if i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			j := i + 1
			for j < len(nums) && nums[j]-nums[j-1] == 1 {
				j++
			}
			res[len(res)-1] += "->" + strconv.Itoa(nums[j-1])
			i = j - 1
		}
	}
	return res
}
