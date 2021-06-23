package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	n := len(beginWord)
	set := make(map[string]bool)
	dist := make(map[string]int)
	dist[beginWord] = 0
	paths := [][]string{}
	path := []string{}
	for _, word := range wordList {
		set[word] = true
	}
	var dfs = func(word string) {
		if word == beginWord {
			p := make([]string, n)
			for i := 0; i < n; i++ {
				p[i] = path[n-1-i]
			}
			paths = append(path, p)
		} else {
			t := []byte(word)
			for i := 0; i < n; i++ {
				for c := 'a'; c <= 'z'; c++ {
					t[i] = byte(c)
					s := string(t)
					if set[s] && dist[s] == 0 {
						dist[s] = dist[word] + 1
					}
				}
			}
		}
	}
	return paths
}
