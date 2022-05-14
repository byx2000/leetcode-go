package maximal_rectangle

// https://leetcode.cn/problems/maximal-rectangle/

func maximalRectangle(nums [][]byte) int {
	if len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 求柱状图中最大矩形面积
	// 参考https://leetcode.cn/problems/largest-rectangle-in-histogram/
	maxAreaOfRow := func(nums []int) int {
		// left[i]表示nums[i]左边第一个比nums[i]小的元素下标
		left := make([]int, len(nums))
		left[0] = -1
		for i := 1; i < len(nums); i++ {
			idx := i - 1
			for idx != -1 && nums[idx] >= nums[i] {
				idx = left[idx]
			}
			left[i] = idx
		}

		// right[i]表示nums[i]右边第一个比nums[i]小的元素下标
		right := make([]int, len(nums))
		right[len(nums)-1] = len(nums)
		for i := len(nums) - 2; i >= 0; i-- {
			idx := i + 1
			for idx != len(nums) && nums[idx] >= nums[i] {
				idx = right[idx]
			}
			right[i] = idx
		}

		maxArea := 0
		for i := 0; i < len(nums); i++ {
			maxArea = max(maxArea, nums[i]*(right[i]-left[i]-1))
		}

		return maxArea
	}

	// 把每一行转换成柱状图
	row := make([]int, len(nums[0]))
	maxArea := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == '0' {
				row[j] = 0
			} else {
				row[j]++
			}
		}
		maxArea = max(maxArea, maxAreaOfRow(row))
	}

	return maxArea
}
