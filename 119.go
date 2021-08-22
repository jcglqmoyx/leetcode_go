package main

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0], res[rowIndex] = 1, 1
	for i := 1; i < rowIndex; i++ {
		res[i] = res[i-1] * (rowIndex - i + 1) / i
	}
	return res
}
