package max_consecutive_ones_iii

import (
	"math"
)

// https://leetcode-cn.com/problems/max-consecutive-ones-iii/

func longestOnes1(nums []int, k int) int {
	i, j := 0, 0
	zeroCnt := 0
	maxLen := 0

	for j < len(nums) {
		if zeroCnt <= k {
			maxLen = int(math.Max(float64(maxLen), float64(j-i)))
			if nums[j] == 0 {
				zeroCnt++
			}
			j++
		} else {
			if nums[i] == 0 {
				zeroCnt--
			}
			i++
		}
	}

	if zeroCnt <= k {
		maxLen = int(math.Max(float64(maxLen), float64(j-i)))
	}

	return maxLen
}

func longestOnes2(nums []int, k int) int {
	i, j := 0, 0
	zeroCnt := 0
	maxLen := 0

	// [i, j)区间内0的个数<=k
	for j < len(nums) {
		// j向右移动，直到zeroCnt > k
		for j < len(nums) && zeroCnt <= k {
			maxLen = int(math.Max(float64(maxLen), float64(j-i)))
			if nums[j] == 0 {
				zeroCnt++
			}
			j++
		}

		// i向右移动，直到zeroCnt <= k
		for zeroCnt > k {
			if nums[i] == 0 {
				zeroCnt--
			}
			i++
		}
	}

	maxLen = int(math.Max(float64(maxLen), float64(j-i)))

	return maxLen
}
