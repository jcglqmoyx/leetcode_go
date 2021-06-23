package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	n := len(beginWord)
	set, dist := make(map[string]bool), make(map[string]int)
	dist[beginWord] = 0
	var paths [][]string
	path := []string{endWord}
	for _, word := range wordList {
		set[word] = true
	}
	var dfs = func(word string) {}
	dfs = func(word string) {
		if word == beginWord {
			p := make([]string, len(path))
			for i := 0; i < len(path); i++ {
				p[i] = path[len(path)-1-i]
			}
			paths = append(paths, p)
		} else {
			t := []byte(word)
			for i := 0; i < n; i++ {
				for c := 'a'; c <= 'z'; c++ {
					tmp := t[i]
					t[i] = byte(c)
					s := string(t)
					t[i] = tmp
					if _, ok := dist[s]; ok && dist[s]+1 == dist[word] {
						path = append(path, s)
						dfs(s)
						path = path[:len(path)-1]
					}
				}
			}
		}
	}
	q := []string{beginWord}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		chars := []byte(t)
		for i := 0; i < len(t); i++ {
			for c := 'a'; c <= 'z'; c++ {
				tmp := chars[i]
				chars[i] = byte(c)
				s := string(chars)
				chars[i] = tmp
				_, ok := dist[s]
				if set[s] && !ok {
					dist[s] = dist[t] + 1
					if s == endWord {
						break
					}
					q = append(q, s)
				}
			}
		}
	}
	dfs(endWord)
	return paths
}
