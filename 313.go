package main

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	f := make([]int, n)
	pointers := make([]int, len(primes))
	f[0] = 1
	for i := 1; i < n; i++ {
		minValue := math.MaxInt32
		for j := 0; j < len(primes); j++ {
			tmp := f[pointers[j]] * primes[j]
			if tmp < minValue {
				minValue = tmp
			}
		}
		f[i] = minValue
		for j := 0; j < len(primes); j++ {
			if f[pointers[j]]*primes[j] == minValue {
				pointers[j]++
			}
		}
	}
	return f[n-1]
}
