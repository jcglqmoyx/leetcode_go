package main

func longestCommonPrefix(strs []string) string {
	minLength := 100000
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	if minLength == 0 {
		return ""
	}
	for i := 0; i < minLength; i++ {
		c := strs[0][i]
		for _, s := range strs {
			if s[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLength]
}
