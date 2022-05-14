package trapping_rain_water

// https://leetcode.cn/problems/trapping-rain-water/

func trap(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// left[i]表示nums[i]左边最高元素的高度
	left := make([]int, len(nums))
	left[0] = 0
	for i := 1; i < len(nums); i++ {
		left[i] = max(nums[i-1], left[i-1])
	}

	// right[i]表示nums[i]右边最高元素的高度
	right := make([]int, len(nums))
	right[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = max(nums[i+1], right[i+1])
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		x := min(left[i], right[i]) - nums[i]
		if x > 0 {
			result += x
		}
	}

	return result
}
