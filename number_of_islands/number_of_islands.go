package number_of_islands

// https://leetcode-cn.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return
		}

		if grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(i, j)
			}
		}
	}

	return cnt
}
