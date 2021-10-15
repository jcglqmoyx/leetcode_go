package main

func longestPalindrome(s string) int {
	cnt := make([]int, 125)
	for _, c := range s {
		cnt[c]++
	}
	res := 0
	hasOdd := false
	for i := 65; i <= 122; i++ {
		if cnt[i]%2 == 0 {
			res += cnt[i]
		} else {
			hasOdd = true
			res += cnt[i] - 1
		}
	}
	if hasOdd {
		res++
	}
	return res
}
