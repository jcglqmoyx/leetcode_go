package main

func numWays(n int, relation [][]int, k int) int {
	f := make([]int, n)
	f[0] = 1
	for i := 0; i < k; i++ {
		cur := make([]int, n)
		for _, r := range relation {
			src, dst := r[0], r[1]
			cur[dst] += f[src]
		}
		copy(f, cur)
	}
	return f[n-1]
}
