package unique_paths_ii

// https://leetcode-cn.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(grid [][]int) int {
	cache := make([][]int, len(grid))
	mark := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		cache[i] = make([]int, len(grid[i]))
		mark[i] = make([]bool, len(grid[i]))
	}

	// dp(r, c)表示从(0, 0)到(r, c)的路径总数
	var dp func(r, c int) int
	dp = func(r, c int) (ret int) {
		if mark[r][c] {
			return cache[r][c]
		}

		defer func() {
			mark[r][c] = true
			cache[r][c] = ret
		}()

		if grid[r][c] == 1 {
			return 0
		} else if r == 0 && c == 0 {
			return 1
		} else if r == 0 {
			return dp(r, c-1)
		} else if c == 0 {
			return dp(r-1, c)
		} else {
			return dp(r, c-1) + dp(r-1, c)
		}
	}

	return dp(len(grid)-1, len(grid[0])-1)
}
