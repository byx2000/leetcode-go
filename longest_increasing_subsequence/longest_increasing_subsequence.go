package longest_increasing_subsequence

import "math"

// https://leetcode.cn/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	cache := make([]int, len(nums))
	mark := make([]bool, len(nums))

	// dp(index)表示以nums[index]结尾的最长递增子序列长度
	var dp func(index int) int
	dp = func(index int) (ret int) {
		if mark[index] {
			return cache[index]
		}

		defer func() {
			mark[index] = true
			cache[index] = ret
		}()

		result := 1
		for i := 0; i < index; i++ {
			if nums[index] > nums[i] {
				result = int(math.Max(float64(result), float64(1+dp(i))))
			}
		}
		return result
	}

	maxLen := 1
	for i := 0; i < len(nums); i++ {
		maxLen = int(math.Max(float64(maxLen), float64(dp(i))))
	}

	return maxLen
}
