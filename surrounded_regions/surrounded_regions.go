package surrounded_regions

// https://leetcode-cn.com/problems/surrounded-regions/

func solve(board [][]byte) {
	flag := make([][]bool, len(board))
	for i := 0; i < len(flag); i++ {
		flag[i] = make([]bool, len(board[i]))
	}

	// 标记所有处于边界的O
	var mark func(r, c int)
	mark = func(r, c int) {
		if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) {
			return
		}

		if flag[r][c] || board[r][c] != 'O' {
			return
		}

		flag[r][c] = true
		mark(r-1, c)
		mark(r+1, c)
		mark(r, c-1)
		mark(r, c+1)
	}

	for i := 0; i < len(board[0]); i++ {
		mark(0, i)
		mark(len(board)-1, i)
	}

	for i := 0; i < len(board); i++ {
		mark(i, 0)
		mark(i, len(board[i])-1)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' && !flag[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
