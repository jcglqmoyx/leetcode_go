package main

func rearrangeSticks(n int, k int) int {
	MOD := int(1e9 + 7)
	f := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = make([]int, k+1)
	}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			f[i][j] = (f[i-1][j-1] + (i-1)*f[i-1][j]) % MOD
		}
	}
	return f[n][k]
}
