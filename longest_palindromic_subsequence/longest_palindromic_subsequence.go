package longest_palindromic_subsequence

import "math"

// https://leetcode.cn/problems/longest-palindromic-subsequence/

func longestPalindromeSubseq(s string) int {
	cache := make([][]int, len(s))
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]int, len(s))
		mark[i] = make([]bool, len(s))
	}

	// dp(i, j)表示s[i~j]的最长回文子序列长度
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i == j {
			return 1
		}

		if i+1 == j {
			if s[i] == s[j] {
				return 2
			}
			return 1
		}

		if s[i] == s[j] {
			return 2 + dp(i+1, j-1)
		}

		return int(math.Max(float64(dp(i+1, j)), float64(dp(i, j-1))))
	}

	return dp(0, len(s)-1)
}
