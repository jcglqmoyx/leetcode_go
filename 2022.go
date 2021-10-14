package main

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	idx := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = original[idx]
			idx++
		}
	}
	return res
}
