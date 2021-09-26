package main

func countBits(n int) []int {
	cnt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cnt[i] = cnt[i>>1]
		if i&1 == 1 {
			cnt[i]++
		}
	}
	return cnt
}
