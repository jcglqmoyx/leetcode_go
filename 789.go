package main

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
func dist(s, t []int) int {
	return abs(s[0]-t[0]) + abs(s[1]-t[1])
}
func escapeGhosts(ghosts [][]int, target []int) bool {
	start := []int{0, 0}
	d := dist(start, target)
	for _, ghost := range ghosts {
		if dist(ghost, target) <= d {
			return false
		}
	}
	return true
}
