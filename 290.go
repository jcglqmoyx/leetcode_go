package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	ps := map[byte]string{}
	sp := map[string]byte{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		ch, word := pattern[i], words[i]
		if sp[word] > 0 && sp[word] != ch || ps[ch] != "" && ps[ch] != word {
			return false
		}
		sp[word] = ch
		ps[ch] = word
	}
	return true
}
