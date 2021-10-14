package main

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l < r {
		mid := l + (r-l+1)/2
		if int64(mid)*int64(mid) <= int64(num) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return int64(l)*int64(l) == int64(num)
}
