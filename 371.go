package main

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	carry, sum := (a&b)<<1, a^b
	return getSum(carry, sum)
}
