package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r + 1) / 2
		if arr[mid] > arr[mid-1] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
