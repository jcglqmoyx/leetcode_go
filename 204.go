package main

func countPrimes(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			isPrime[i] = true
			primes = append(primes, i)
		}
		for j := 0; primes[j]*i < n; j++ {
			isPrime[primes[j]*i] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
	return len(primes)
}
