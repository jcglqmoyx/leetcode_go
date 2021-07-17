package main

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := l + (r-l+1)/2
		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
