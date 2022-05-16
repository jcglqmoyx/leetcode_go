package main

func corpFlightBookings(bookings [][]int, n int) []int {
	d := make([]int, n)
	for _, b := range bookings {
		d[b[0]-1] += b[2]
		if b[1] < n {
			d[b[1]] -= b[2]
		}
	}
	for i := 1; i < n; i++ {
		d[i] += d[i-1]
	}
	return d
}
