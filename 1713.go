package main

import "sort"

func minOperations(target []int, arr []int) int {
	hash := make(map[int]int, len(target))
	for i := 0; i < len(target); i++ {
		hash[target[i]] = i
	}
	var seq []int
	for _, num := range arr {
		if idx, ok := hash[num]; ok {
			if len(seq) == 0 || seq[len(seq)-1] < idx {
				seq = append(seq, idx)
			} else {
				pos := sort.SearchInts(seq, idx)
				seq[pos] = idx
			}
		}
	}
	return len(target) - len(seq)
}
