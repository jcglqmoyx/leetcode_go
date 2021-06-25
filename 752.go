package main

func openLock(deadends []string, target string) int {
	start := "0000"
	if target == start {
		return 0
	}
	dead := map[string]bool{}
	for _, d := range deadends {
		dead[d] = true
	}
	if dead[start] {
		return -1
	}
	dist := map[string]int{}
	dist[start] = 0
	q := []string{start}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		for i := 0; i < len(t); i++ {
			for j := -1; j <= 1; j += 2 {
				ch := t[i]
				ch = byte((int(ch)-int('0')+j+10)%10 + int('0'))
				s := t[:i] + string(ch) + t[i+1:]
				if _, ok := dist[s]; !ok && !dead[s] {
					dist[s] = dist[t] + 1
					if s == target {
						return dist[s]
					}
					q = append(q, s)
				}
			}
		}
	}
	return -1
}
