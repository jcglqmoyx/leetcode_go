package main

import "strconv"

func compress(chars []byte) int {
	i := 0
	for j, k := 0, 0; j < len(chars); j++ {
		chars[i] = chars[j]
		i++
		k = j
		for k < len(chars) && chars[k] == chars[j] {
			k++
		}
		cnt := k - j
		if cnt > 1 {
			s := strconv.Itoa(cnt)
			for _, ch := range s {
				chars[i] = byte(ch)
				i++
			}
		}
		j = k - 1
	}
	return i
}
