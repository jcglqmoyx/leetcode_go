package main

import (
	"math/rand"
	"sort"
	"time"
)

type pair struct{ pow, idx int }

func kWeakestRows(mat [][]int, k int) []int {
	m := len(mat)
	pairs := make([]pair, m)
	for i, row := range mat {
		pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		pairs[i] = pair{pow, i}
	}
	rand.Seed(time.Now().UnixNano())
	randomizedSelected(pairs, 0, m-1, k)
	pairs = pairs[:k]
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
	})
	ans := make([]int, k)
	for i, p := range pairs {
		ans[i] = p.idx
	}
	return ans
}

func randomizedSelected(a []pair, l, r, k int) {
	if l >= r {
		return
	}
	pos := randomPartition(a, l, r)
	num := pos - l + 1
	if k == num {
		return
	}
	if k < num {
		randomizedSelected(a, l, pos-1, k)
	} else {
		randomizedSelected(a, pos+1, r, k-num)
	}
}

func randomPartition(a []pair, l, r int) int {
	i := rand.Intn(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []pair, l, r int) int {
	pivot := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j].pow < pivot.pow || a[j].pow == pivot.pow && a[j].idx <= pivot.idx {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
