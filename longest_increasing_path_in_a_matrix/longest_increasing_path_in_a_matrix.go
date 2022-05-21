package longest_increasing_path_in_a_matrix

// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/

func longestIncreasingPath(matrix [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	cache := make([][]int, len(matrix))
	mark := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		cache[i] = make([]int, len(matrix[i]))
		mark[i] = make([]bool, len(matrix[i]))
	}

	// 以matrix[i][j]结尾的最长递增路径长度
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		result := 1
		if i > 0 && matrix[i][j] > matrix[i-1][j] {
			result = max(result, 1+dp(i-1, j))
		}
		if i < len(matrix)-1 && matrix[i][j] > matrix[i+1][j] {
			result = max(result, 1+dp(i+1, j))
		}
		if j > 0 && matrix[i][j] > matrix[i][j-1] {
			result = max(result, 1+dp(i, j-1))
		}
		if j < len(matrix[i])-1 && matrix[i][j] > matrix[i][j+1] {
			result = max(result, 1+dp(i, j+1))
		}

		return result
	}

	maxLen := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxLen = max(maxLen, dp(i, j))
		}
	}

	return maxLen
}
