package nearest_exit_from_entrance_in_maze

// https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/

func nearestExit(maze [][]byte, entrance []int) int {
	// 走过的路用'*'标记
	queue := [][]int{entrance}
	maze[entrance[0]][entrance[1]] = '*'

	isExit := func(r, c int) bool {
		return (r == 0 || r == len(maze)-1 || c == 0 || c == len(maze[r])-1) &&
			!(r == entrance[0] && c == entrance[1])
	}

	addQueue := func(r, c int) {
		if r < 0 || r >= len(maze) || c < 0 || c >= len(maze[r]) || maze[r][c] != '.' {
			return
		}
		maze[r][c] = '*'
		queue = append(queue, []int{r, c})
	}

	dis := 0
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			cur := queue[0]
			queue = queue[1:]
			if isExit(cur[0], cur[1]) {
				return dis
			}

			addQueue(cur[0]-1, cur[1])
			addQueue(cur[0]+1, cur[1])
			addQueue(cur[0], cur[1]-1)
			addQueue(cur[0], cur[1]+1)
		}
		dis++
	}

	return -1
}
