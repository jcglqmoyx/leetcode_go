package main

import "math"

func arrangeCoins(n int) int {
	return (int(math.Sqrt(float64(1+n*8))) - 1) / 2
}
