package main

func maxProfit(prices []int) int {
	var min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	minPrice, res := 1000000, 0
	for _, p := range prices {
		minPrice = min(minPrice, p)
		res = max(res, p-minPrice)
	}
	return res
}
