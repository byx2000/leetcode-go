package corporate_flight_bookings

// https://leetcode.cn/problems/corporate-flight-bookings/

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for _, b := range bookings {
		diff[b[0]-1] += b[2]
		if b[1] < len(diff) {
			diff[b[1]] -= b[2]
		}
	}

	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]
	}

	return diff
}
