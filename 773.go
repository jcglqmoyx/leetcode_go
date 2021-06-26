package main

import (
	"strconv"
)

func slidingPuzzle(board [][]int) int {
	var getStatus = func(a [][]int) string {
		var res string
		for _, i := range a {
			for _, j := range i {
				res += strconv.Itoa(j)
			}
		}
		return res
	}
	target := "123450"
	if getStatus(board) == target {
		return 0
	}
	var dist = make(map[string]int)
	dist[getStatus(board)] = 0
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	var q [][][]int
	q = append(q, board)
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		x, y := -1, -1
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				if t[i][j] == 0 {
					x = i
					y = j
				}
			}
			for i := 0; i < 4; i++ {
				a, b := x+dx[i], y+dy[i]
				if a < 0 || a == 2 || b < 0 || b == 3 {
					continue
				}
				var r = make([][]int, 2)
				for row := 0; row < 2; row++ {
					r[row] = make([]int, 3)
					for col := 0; col < 3; col++ {
						r[row][col] = t[row][col]
					}
				}
				r[x][y], r[a][b] = r[a][b], r[x][y]

				status := getStatus(r)
				if _, ok := dist[status]; !ok {
					dist[status] = dist[getStatus(t)] + 1
					if status == target {
						return dist[status]
					}
					q = append(q, r)
				}
			}
		}
	}
	return -1
}
