package main

func findErrorNums(nums []int) []int {
	hash := make([]int, len(nums))
	for _, num := range nums {
		hash[num-1]++
	}
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if hash[i] == 2 {
			a = i + 1
		} else if hash[i] == 0 {
			b = i + 1
		}
	}
	return []int{a, b}
}
