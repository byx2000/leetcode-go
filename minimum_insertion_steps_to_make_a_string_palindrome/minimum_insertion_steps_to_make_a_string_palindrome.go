package minimum_insertion_steps_to_make_a_string_palindrome

import "math"

// https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/

func minInsertions(s string) int {
	cache := make([][]int, len(s))
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]int, len(s))
		mark[i] = make([]bool, len(s))
	}

	// dp(i, j)表示将s[i~j]转换成回文串的最少操作数
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i >= j {
			return 0
		}

		if s[i] == s[j] {
			return dp(i+1, j-1)
		}

		return int(1 + math.Min(float64(dp(i+1, j)), float64(dp(i, j-1))))
	}

	return dp(0, len(s)-1)
}
