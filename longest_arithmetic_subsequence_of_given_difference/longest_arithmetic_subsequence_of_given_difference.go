package longest_arithmetic_subsequence_of_given_difference

import "math"

// https://leetcode.cn/problems/longest-arithmetic-subsequence-of-given-difference/

func longestSubsequence1(nums []int, d int) int {
	cache := make([]int, len(nums))
	mark := make([]bool, len(nums))

	// dp(index)表示以nums[index]结尾的最长等差子序列的长度
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
			if nums[index]-nums[i] == d {
				result = int(math.Max(float64(result), float64(dp(i)+1)))
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

func longestSubsequence2(nums []int, d int) int {
	dp := map[int]int{}
	maxLen := 1

	for _, n := range nums {
		if v, exist := dp[n-d]; exist {
			dp[n] = v + 1
		} else {
			dp[n] = 1
		}
		maxLen = int(math.Max(float64(maxLen), float64(dp[n])))
	}

	return maxLen
}
