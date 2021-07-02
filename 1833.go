package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	cnt := 0
	sort.Ints(costs)
	for _, cost := range costs {
		if coins >= cost {
			cnt++
			coins -= cost
		} else {
			break
		}
	}
	return cnt
}
