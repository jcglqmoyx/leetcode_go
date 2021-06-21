package main

func romanToInt(s string) int {
	var hash = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	for i := 0; i < len(s); i++ {
		res += hash[s[i]]
		for j := i - 1; j >= 0; j-- {
			if hash[s[j]] < hash[s[i]] {
				res -= 2 * hash[s[j]]
			}
		}
	}
	return res
}
