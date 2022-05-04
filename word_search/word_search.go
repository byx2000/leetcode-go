package word_search

// https://leetcode-cn.com/problems/word-search/

func exist(board [][]byte, word string) bool {
	var visit [][]bool
	for i := 0; i < len(board); i++ {
		visit = append(visit, []bool{})
		for range board[i] {
			visit[i] = append(visit[i], false)
		}
	}

	var dfs func(index, r, c int) bool
	dfs = func(index, r, c int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) {
			return false
		}

		if visit[r][c] {
			return false
		}

		visit[r][c] = true

		if board[r][c] != word[index] {
			visit[r][c] = false
			return false
		}

		result := dfs(index+1, r-1, c) ||
			dfs(index+1, r+1, c) ||
			dfs(index+1, r, c-1) ||
			dfs(index+1, r, c+1)

		visit[r][c] = false
		return result
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}

	return false
}
