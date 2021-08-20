package main

func convertToTitle(columnNumber int) string {
	str := ""
	for columnNumber > 0 {
		columnNumber--
		str += string('A' + byte(columnNumber%26))
		columnNumber /= 26
	}
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}
