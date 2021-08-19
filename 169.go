package main

func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for _, x := range nums {
		if x == res {
			cnt++
		} else {
			cnt--
			if cnt < 0 {
				cnt = 0
				res = x
			}
		}
	}
	return res
}
