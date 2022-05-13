package maximal_square

import "math"

// https://leetcode.cn/problems/maximal-square/

func maximalSquare(nums [][]byte) int {
	var min = func(a int, args ...int) int {
		r := a
		for _, n := range args {
			if n < r {
				r = n
			}
		}
		return r
	}

	cache := make([][]int, len(nums))
	mark := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]int, len(nums[i]))
		mark[i] = make([]bool, len(nums[i]))
	}

	// dp(i, j)表示以nums[i][j]为右下角的最大正方形边长
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if i < 0 || j < 0 {
			return 0
		}

		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if nums[i][j] == '0' {
			return 0
		}

		return 1 + min(dp(i-1, j), dp(i, j-1), dp(i-1, j-1))
	}

	maxLen := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == '1' {
				maxLen = int(math.Max(float64(maxLen), float64(dp(i, j))))
			}
		}
	}

	return maxLen * maxLen
}
