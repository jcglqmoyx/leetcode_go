package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	res := ""
	carry, sum := 0, 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			sum = int(a[i]-'0') + int(b[j]-'0') + carry
			i--
			j--
		} else if i >= 0 {
			sum = int(a[i]-'0') + carry
			i--
		} else {
			sum = int(b[j]-'0') + carry
			j--
		}
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}
