package main

//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
//
//func reverseStr(s string, k int) string {
//	chars := []byte(s)
//	for i := 0; i < len(chars); i++ {
//		l, r := i, min(i+k-1, len(chars)-1)
//		for l < r {
//			chars[l], chars[r] = chars[r], chars[l]
//			l++
//			r--
//		}
//		i = i + k*2 - 1
//	}
//	return string(chars)
//}
