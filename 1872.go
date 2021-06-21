package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stoneGameVIII(stones []int) int {
	n := len(stones)
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			sum[i] = stones[i]
		} else {
			sum[i] = sum[i-1] + stones[i]
		}
	}
	f := make([]int, n)
	f[n-1] = sum[n-1]
	for i := n - 2; i > 0; i-- {
		f[i] = max(f[i+1], sum[i]-f[i+1])
	}
	return f[1]
}
