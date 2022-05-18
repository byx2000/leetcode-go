package shortest_bridge

// https://leetcode.cn/problems/shortest-bridge/

func shortestBridge(grid [][]int) int {
	var queue [][]int

	// 将某个岛的所有1变成2，并将该岛周围的所有0变成3
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] == 2 || grid[r][c] == 3 {
			return
		}

		if grid[r][c] == 0 {
			queue = append(queue, []int{r, c})
			grid[r][c] = 3

			return
		}

		grid[r][c] = 2
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

loop:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				break loop
			}
		}
	}

	addQueue := func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return
		}
		if grid[r][c] != 0 && grid[r][c] != 1 {
			return
		}
		if grid[r][c] == 0 {
			grid[r][c] = 3
		}
		queue = append(queue, []int{r, c})
	}

	// BFS扩展该岛周围的0
	dis := 0
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			cur := queue[0]
			queue = queue[1:]
			if grid[cur[0]][cur[1]] == 1 {
				return dis
			}

			addQueue(cur[0]+1, cur[1])
			addQueue(cur[0]-1, cur[1])
			addQueue(cur[0], cur[1]+1)
			addQueue(cur[0], cur[1]-1)
		}
		dis++
	}

	return dis
}
