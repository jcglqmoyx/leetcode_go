package main

func thirdMax(nums []int) int {
	var first, second, third = -2 << 32, -2 << 32, -2 << 32
	for _, num := range nums {
		if num > third {
			if num > second {
				if num > first {
					third = second
					second = first
					first = num
				} else if num < first {
					third = second
					second = num
				}
			} else if num < second {
				third = num
			}
		}
	}
	if third == -2<<32 {
		return first
	}
	return third
}
