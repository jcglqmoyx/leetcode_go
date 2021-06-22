package main

func maxDistance(grid [][]int) int {
	q := make([][]int, 0)
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	if len(q) == n*m || len(q) == 0 {
		return -1
	}
	res := 0
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for len(q) != 0 {
		flag := false
		for size := len(q); size > 0; size-- {
			t := q[0]
			q = q[1:]
			for i := 0; i < 4; i++ {
				x := t[0] + dx[i]
				y := t[1] + dy[i]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 0 {
					q = append(q, []int{x, y})
					grid[x][y] = 1
					flag = true
				}
			}
		}
		if flag {
			res++
		}
	}
	return res
}
