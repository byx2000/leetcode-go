package max_area_of_island

import "math"

// https://leetcode-cn.com/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) int {
	var area func(r, c int) int
	area = func(r, c int) int {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return 0
		}

		if grid[r][c] == 0 {
			return 0
		}

		grid[r][c] = 0
		return 1 + area(r-1, c) + area(r+1, c) + area(r, c-1) + area(r, c+1)
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				maxArea = int(math.Max(float64(maxArea), float64(area(i, j))))
			}
		}
	}

	return maxArea
}
