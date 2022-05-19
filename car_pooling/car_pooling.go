package car_pooling

// https://leetcode.cn/problems/car-pooling/

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for _, t := range trips {
		diff[t[1]] += t[0]
		diff[t[2]] -= t[0]
	}

	sum := 0
	for i := 0; i < len(diff); i++ {
		sum += diff[i]
		if sum > capacity {
			return false
		}
	}

	return true
}
