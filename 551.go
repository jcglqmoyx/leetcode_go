package main

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func checkRecord(s string) bool {
	absent := false
	for i, c := range s {
		if c == 'A' {
			if absent {
				return false
			}
			absent = true
		} else if c == 'L' {
			cntLate := 1
			for j := i + 1; j <= min(i+2, len(s)-1); j++ {
				if s[j] == 'L' {
					cntLate++
				}
			}
			if cntLate == 3 {
				return false
			}
		}
	}
	return true
}
