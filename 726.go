package main

import (
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	u := 0
	var dfs = func(s string) map[string]int {
		return nil
	}
	dfs = func(s string) map[string]int {
		res := make(map[string]int)
		for u < len(s) {
			if s[u] == '(' {
				u++
				t := dfs(s)
				u++
				k := u
				for k < len(s) && unicode.IsDigit(rune(s[k])) {
					k++
				}
				cnt := 1
				if k > u {
					cnt, _ = strconv.Atoi(s[u:k])
					u = k
				}
				for atom, number := range t {
					res[atom] += number * cnt
				}
			} else if s[u] == ')' {
				break
			} else {
				cnt, k := 1, u+1
				for k < len(s) && unicode.IsLower(rune(s[k])) {
					k++
				}
				atom := s[u:k]
				u = k
				for k < len(s) && unicode.IsDigit(rune(s[k])) {
					k++
				}
				if k > u {
					cnt, _ = strconv.Atoi(s[u:k])
				}
				u = k
				res[atom] += cnt
			}
		}
		return res
	}
	res := ""
	var arr []string
	t := dfs(formula)
	for k := range t {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	for _, atom := range arr {
		res += atom
		if t[atom] > 1 {
			res += strconv.Itoa(t[atom])
		}
	}
	return res
}
