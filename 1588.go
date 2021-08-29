package main

func sumOddLengthSubarrays(arr []int) int {
	res, n := 0, len(arr)
	for i := 0; i < n; i++ {
		leftOdd, rightOdd, leftEven, rightEven := (i+1)/2, (n-i)/2, (i+2)/2, (n-i+1)/2
		res += arr[i] * (leftOdd*rightOdd + leftEven*rightEven)
	}
	return res
}
