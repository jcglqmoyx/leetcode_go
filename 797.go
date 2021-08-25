package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var path []int
	var paths [][]int
	path = append(path, 0)
	var dfs func(int)
	dfs = func(start int) {
		if start == len(graph)-1 {
			paths = append(paths, append([]int(nil), path...))
			return
		}
		for _, ne := range graph[start] {
			path = append(path, ne)
			dfs(ne)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return paths
}
