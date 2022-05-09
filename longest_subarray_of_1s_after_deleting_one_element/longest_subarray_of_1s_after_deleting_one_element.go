package longest_subarray_of_1s_after_deleting_one_element

import "math"

// https://leetcode-cn.com/problems/longest-subarray-of-1s-after-deleting-one-element/

func longestSubarray(nums []int) int {
	i, j := 0, 0
	zeroCnt := 0
	maxLen := 0

	// 区间[i, j)中0的个数 <= 1
	for j < len(nums) {
		// 向右移动j，直到zeroCnt > 1
		for j < len(nums) && zeroCnt <= 1 {
			maxLen = int(math.Max(float64(maxLen), float64(j-i-1)))
			if nums[j] == 0 {
				zeroCnt++
			}
			j++
		}

		// 向右移动i，直到zeroCnt <= 1
		for zeroCnt > 1 {
			if nums[i] == 0 {
				zeroCnt--
			}
			i++
		}
	}

	maxLen = int(math.Max(float64(maxLen), float64(j-i-1)))

	return maxLen
}
