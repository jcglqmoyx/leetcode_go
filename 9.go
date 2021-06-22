package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x > 0 && x%10 == 0 {
		return false
	}
	var y = 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	return y/10 == x || y == x
}
