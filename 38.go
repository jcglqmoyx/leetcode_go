package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := "1"
	for i := 2; i <= n; i++ {
		t := ""
		for j := 0; j < len(s); j++ {
			k := j + 1
			for k < len(s) {
				if s[k] == s[k-1] {
					k++
				} else {
					break
				}
			}
			cnt := k - j
			j = k - 1
			t += strconv.Itoa(cnt) + string(s[j])
		}
		s = t
	}
	return s
}
