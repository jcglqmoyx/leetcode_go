package main

func findTheDifference(s string, t string) byte {
	m := map[byte]int{}
	for _, ch := range s {
		m[byte(ch)]++
	}
	for _, ch := range t {
		m[byte(ch)]--
	}
	for k, v := range m {
		if v == -1 {
			return k
		}
	}
	return 'a'
}
