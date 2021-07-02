package main

import "sort"

func frequencySort(s string) string {
	cnt := make(map[byte]int)
	var chars []byte
	for _, c := range s {
		_, ok := cnt[byte(c)]
		if !ok {
			chars = append(chars, byte(c))
		}
		cnt[byte(c)]++
	}
	sort.Slice(chars, func(i, j int) bool {
		return cnt[chars[i]] > cnt[chars[j]]
	})
	var res []byte
	for _, value := range chars {
		num := cnt[value]
		for i := 0; i < num; i++ {
			res = append(res, value)
		}
	}
	return string(res)
}
