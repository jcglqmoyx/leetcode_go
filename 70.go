package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first, second, third := 1, 2, 0
	for i := 3; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}
	return third
}
