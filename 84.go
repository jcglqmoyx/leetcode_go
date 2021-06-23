package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	var stk []int
	for i := 0; i < n; i++ {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			left[i] = -1
		} else {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			right[i] = n
		} else {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	res := 0
	for i := 0; i < n; i++ {
		area := heights[i] * (right[i] - left[i] - 1)
		if area > res {
			res = area
		}
	}
	return res
}
