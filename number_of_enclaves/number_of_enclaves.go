package number_of_enclaves

// https://leetcode-cn.com/problems/number-of-enclaves/

func numEnclaves(grid [][]int) int {
	// 将相连的1全部变成0，并返回1的数量
	var mark func(r, c int)
	mark = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return
		}

		if grid[r][c] == 0 {
			return
		}

		grid[r][c] = 0
		mark(r-1, c)
		mark(r+1, c)
		mark(r, c-1)
		mark(r, c+1)
	}

	// 标记边界的1
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 1 {
			mark(0, i)
		}
		if grid[len(grid)-1][i] == 1 {
			mark(len(grid)-1, i)
		}
	}
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
			mark(i, 0)
		}
		if grid[i][len(grid[0])-1] == 1 {
			mark(i, len(grid[0])-1)
		}
	}

	// 计数
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}
	}

	return cnt
}
