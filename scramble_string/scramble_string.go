package scramble_string

// https://leetcode.cn/problems/scramble-string/

func isScramble1(s1 string, s2 string) bool {
	cache := make([][][][]bool, len(s1))
	mark := make([][][][]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		cache[i] = make([][][]bool, len(s1))
		mark[i] = make([][][]bool, len(s1))
		for j := 0; j < len(s2); j++ {
			cache[i][j] = make([][]bool, len(s2))
			mark[i][j] = make([][]bool, len(s2))
			for k := 0; k < len(s2); k++ {
				cache[i][j][k] = make([]bool, len(s2))
				mark[i][j][k] = make([]bool, len(s2))
			}
		}
	}

	// s1[i~j]和s2[m~n]是否互为扰乱字符串
	var dp func(i, j, m, n int) bool
	dp = func(i, j, m, n int) (ret bool) {
		if mark[i][j][m][n] {
			return cache[i][j][m][n]
		}

		defer func() {
			mark[i][j][m][n] = true
			cache[i][j][m][n] = ret
		}()

		if i == j && m == n {
			return s1[i] == s2[m]
		}

		for k := 0; k < j-i; k++ {
			if dp(i, i+k, m, m+k) && dp(i+k+1, j, m+k+1, n) {
				return true
			}
			if dp(i, i+k, n-k, n) && dp(i+k+1, j, m, n-k-1) {
				return true
			}
		}

		return false
	}

	return dp(0, len(s1)-1, 0, len(s2)-1)
}

func isScramble2(s1 string, s2 string) bool {
	cache := make([][][]bool, len(s1))
	mark := make([][][]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		cache[i] = make([][]bool, len(s1))
		mark[i] = make([][]bool, len(s1))
		for j := 0; j < len(s1); j++ {
			cache[i][j] = make([]bool, len(s1)+1)
			mark[i][j] = make([]bool, len(s1)+1)
		}
	}

	// s1[i~i+m-1]和s2[j~j+m-1]是否互为扰乱字符串
	var dp func(i, j, m int) bool
	dp = func(i, j, m int) (ret bool) {
		if mark[i][j][m] {
			return cache[i][j][m]
		}

		defer func() {
			mark[i][j][m] = true
			cache[i][j][m] = ret
		}()

		if m == 1 {
			return s1[i] == s2[j]
		}

		for k := 1; k < m; k++ {
			if dp(i, j, k) && dp(i+k, j+k, m-k) {
				return true
			}
			if dp(i, j+m-k, k) && dp(i+k, j, m-k) {
				return true
			}
		}

		return false
	}

	return dp(0, 0, len(s1))
}
