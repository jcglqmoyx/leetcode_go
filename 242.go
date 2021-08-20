package main

func isAnagram(s string, t string) bool {
	m := map[byte]int{}
	for _, v := range s {
		m[byte(v)]++
	}
	for _, v := range t {
		m[byte(v)]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
