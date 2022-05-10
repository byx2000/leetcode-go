package longest_common_subsequence

import "math"

// https://leetcode.cn/problems/longest-common-subsequence/

func longestCommonSubsequence(s1 string, s2 string) int {
	cache := make([][]int, len(s1))
	mark := make([][]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		cache[i] = make([]int, len(s2))
		mark[i] = make([]bool, len(s2))
	}

	// dp(i, j)表示s1[0~i]和s2[0~j]的最长公共子序列长度
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if i < 0 || j < 0 {
			return 0
		}

		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if s1[i] == s2[j] {
			return 1 + dp(i-1, j-1)
		}
		return int(math.Max(float64(dp(i-1, j)), float64(dp(i, j-1))))
	}

	return dp(len(s1)-1, len(s2)-1)
}
