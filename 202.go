package main

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 {
		sum, num := 0, n
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num /= 10
		}
		if m[sum] {
			return false
		}
		m[sum] = true
		n = sum
	}
	return true
}
