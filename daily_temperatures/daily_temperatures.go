package daily_temperatures

// https://leetcode.cn/problems/daily-temperatures/

func dailyTemperatures1(temperatures []int) []int {
	var stack []int
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			topIdx := stack[len(stack)-1]
			result[topIdx] = i - topIdx
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

func dailyTemperatures2(nums []int) []int {
	result := make([]int, len(nums))

	// dp[i]表示nums[i]右边第一个比nums[i]大的元素索引
	dp := make([]int, len(nums))
	dp[len(nums)-1] = len(nums)

	for i := len(nums) - 2; i >= 0; i-- {
		idx := i + 1
		for idx != len(nums) && nums[idx] <= nums[i] {
			idx = dp[idx]
		}
		dp[i] = idx
		if dp[i] == len(nums) {
			result[i] = 0
		} else {
			result[i] = dp[i] - i
		}
	}

	return result
}
