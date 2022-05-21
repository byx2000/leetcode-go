package longest_chunked_palindrome_decomposition

import "math"

// https://leetcode.cn/problems/longest-chunked-palindrome-decomposition/

func longestDecomposition(s string) int {
	cache := make([][]int, len(s))
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]int, len(s))
		mark[i] = make([]bool, len(s))
	}

	// dp(i, j)表示s[i~j]的最大段式回文数量
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i > j {
			return 0
		}

		maxCnt := 1
		for k := 1; k <= (j-i+1)/2; k++ {
			if s[i:i+k] == s[j-k+1:j+1] {
				maxCnt = int(math.Max(float64(maxCnt), float64(dp(i+k, j-k)+2)))
			}
		}

		return maxCnt
	}

	return dp(0, len(s)-1)
}
