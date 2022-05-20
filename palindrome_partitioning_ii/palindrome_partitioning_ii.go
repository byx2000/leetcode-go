package palindrome_partitioning_ii

import "math"

// https://leetcode.cn/problems/palindrome-partitioning-ii/

func minCut(s string) int {
	cache := make([][]bool, len(s))
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]bool, len(s))
		mark[i] = make([]bool, len(s))
	}

	// 判断s[i~j]是不是回文串
	var isPalindromic func(i, j int) bool
	isPalindromic = func(i, j int) (ret bool) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i == j {
			return true
		}

		if i+1 == j {
			return s[i] == s[j]
		}

		return s[i] == s[j] && isPalindromic(i+1, j-1)
	}

	cache2 := make([]int, len(s))
	mark2 := make([]bool, len(s))

	// dp(index)表示将s[0~index]分割成多个回文子串的最少分割次数
	var dp func(index int) int
	dp = func(index int) (ret int) {
		if mark2[index] {
			return cache2[index]
		}

		defer func() {
			mark2[index] = true
			cache2[index] = ret
		}()

		if isPalindromic(0, index) {
			return 0
		}

		result := index + 1
		for i := index - 1; i >= 0; i-- {
			if isPalindromic(i+1, index) {
				result = int(math.Min(float64(result), float64(dp(i)+1)))
			}
		}

		return result
	}

	return dp(len(s) - 1)
}
