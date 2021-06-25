package main

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	ne := make([]int, n)
	ne[0] = -1
	for i, j := 1, -1; i < n; i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = ne[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		ne[i] = j
	}
	for i, j := 0, -1; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = ne[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == n-1 {
			return i - j
		}
	}
	return -1
}
