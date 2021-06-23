package main

import "fmt"

func main() {
	var a [][]string
	b := []string{"hell", "good"}
	a = append(a, b)
	b[1] = "not"
	fmt.Println(a)
}
