package main

func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m {
		for j < n && s[i] != t[j] {
			j++
		}
		if j >= n {
			break
		}
		i++
		j++
	}
	return i == m
}
