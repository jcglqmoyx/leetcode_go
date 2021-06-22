package main

import "math"

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	if negative {
		return -res
	} else {
		return res
	}
}
