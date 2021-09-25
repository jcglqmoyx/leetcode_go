package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	f := make([][]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if byte(word1[i-1]) == byte(word2[j-1]) {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	return n + m - f[n][m]*2
}
