package main

func minimumMoves(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			cnt++
			i += 2
		}
	}
	return cnt
}
