package main

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
func getMinDistance(nums []int, target int, start int) int {
	res := int(1e8)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			dist := abs(i - start)
			if dist < res {
				res = dist
			}
		}
	}
	return res
}
