package main

func maxPoints(points [][]int) int {
	res := 0
	var equals = func(a []int, b []int) bool {
		return a[0] == b[0] && a[1] == b[1]
	}
	var max = func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for _, p := range points {
		ss, vs := 0, 0
		hash := map[float64]int{}
		for _, q := range points {
			if equals(p, q) {
				ss++
			} else if p[0] == q[0] {
				vs++
			} else {
				slope := float64(p[1]-q[1]) / float64(p[0]-q[0])
				hash[slope]++
			}
			cnt := vs
			for _, v := range hash {
				cnt = max(cnt, v)
			}
			res = max(res, cnt+ss)
		}
	}
	return res
}
