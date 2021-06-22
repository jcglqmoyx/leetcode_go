package main

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	f := make([][3]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			f[i][j] = 0x3f3f3f3f
		}
	}
	f[0][0] = 1
	f[0][2] = 1
	f[0][1] = 0
	for i := 1; i < n; i++ {
		if obstacles[i] != 1 {
			f[i][0] = f[i-1][0]
		}
		if obstacles[i] != 2 {
			f[i][1] = f[i-1][1]
		}
		if obstacles[i] != 3 {
			f[i][2] = f[i-1][2]
		}
		if obstacles[i] != 1 {
			f[i][0] = min(f[i][0], min(f[i][1], f[i][2])+1)
		}
		if obstacles[i] != 2 {
			f[i][1] = min(f[i][1], min(f[i][0], f[i][2])+1)
		}
		if obstacles[i] != 3 {
			f[i][2] = min(f[i][2], min(f[i][0], f[i][1])+1)
		}
	}
	return min(f[n-1][0], min(f[n-1][1], f[n-1][2]))
}
