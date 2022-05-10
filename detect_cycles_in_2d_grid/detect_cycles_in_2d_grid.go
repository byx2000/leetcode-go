package detect_cycles_in_2d_grid

// https://leetcode.cn/problems/detect-cycles-in-2d-grid/

func containsCycle(grid [][]byte) bool {
	mark := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mark[i] = make([]bool, len(grid[i]))
	}

	find := false

	var dfs func(pr, pc, r, c int, v byte)
	dfs = func(pr, pc, r, c int, v byte) {
		if find {
			return
		}
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return
		}
		if grid[r][c] != v {
			return
		}
		if mark[r][c] {
			find = true
			return
		}

		mark[r][c] = true
		if r+1 != pr || c != pc {
			dfs(r, c, r+1, c, v)
		}
		if r-1 != pr || c != pc {
			dfs(r, c, r-1, c, v)
		}
		if r != pr || c+1 != pc {
			dfs(r, c, r, c+1, v)
		}
		if r != pr || c-1 != pc {
			dfs(r, c, r, c-1, v)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !mark[i][j] {
				find = false
				dfs(-1, -1, i, j, grid[i][j])
				if find {
					return true
				}
			}
		}
	}

	return false
}
