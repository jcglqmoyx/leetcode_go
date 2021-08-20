package main

func titleToNumber(columnTitle string) int {
	num := 0
	for _, value := range columnTitle {
		num = num*26 + int(value-'A') + 1
	}
	return num
}
