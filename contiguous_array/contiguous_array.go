package contiguous_array

import "math"

// https://leetcode.cn/problems/contiguous-array/

func findMaxLength(nums []int) int {
	// 将0变为-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	// 求和为0的最长子数组长度
	sum := 0
	m := map[int]int{0: -1}
	maxLen := 0
	for i, n := range nums {
		sum += n
		if v, exist := m[sum]; exist {
			maxLen = int(math.Max(float64(maxLen), float64(i-v)))
		} else {
			m[sum] = i
		}
	}

	return maxLen
}
