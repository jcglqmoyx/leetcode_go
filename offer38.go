package main

import "sort"

func permutation(s string) []string {
	var res []string
	var used []bool = make([]bool, len(s))
	var t []byte = []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	var dfs func(cur []byte, index int)
	dfs = func(cur []byte, index int) {
		if index == len(t) {
			res = append(res, string(cur))
			return
		}
		for i := 0; i < len(t); i++ {
			if used[i] || i > 0 && t[i] == t[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			cur = append(cur, t[i])
			dfs(cur, index+1)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	dfs(nil, 0)
	return res
}
