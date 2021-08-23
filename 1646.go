package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}
	arr := make([]int, n+1)
	arr[1] = 1
	res := 1
	for i := 2; i < n+1; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + arr[(i+1)/2]
		}
		res = max(res, arr[i])
	}
	return res
}
