package unique_paths

// https://leetcode-cn.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	cache := make([][]int, m+1)
	mark := make([][]bool, m+1)
	for i := 1; i <= m; i++ {
		cache[i] = make([]int, n+1)
		mark[i] = make([]bool, n+1)
	}

	// dp(r, c)表示从(1, 1)到(r, c)的路径总数
	var dp func(r, c int) int
	dp = func(r, c int) (ret int) {
		if mark[r][c] {
			return cache[r][c]
		}

		defer func() {
			mark[r][c] = true
			cache[r][c] = ret
		}()

		if r == 1 || c == 1 {
			return 1
		} else if r == 1 {
			return dp(r, c-1)
		} else if c == 1 {
			return dp(r-1, c)
		} else {
			return dp(r, c-1) + dp(r-1, c)
		}
	}

	return dp(m, n)
}
