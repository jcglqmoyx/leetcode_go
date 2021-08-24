package main

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	INF := 0x3f3f3f3f
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[src] = 0
	for i := 0; i < k+1; i++ {
		cur := make([]int, n)
		copy(cur, dist)
		for _, flight := range flights {
			a, b, c := flight[0], flight[1], flight[2]
			cur[b] = min(cur[b], dist[a]+c)
		}
		for j := 0; j < n; j++ {
			dist[j] = cur[j]
		}
	}
	if dist[dst] == INF {
		return -1
	}
	return dist[dst]
}
