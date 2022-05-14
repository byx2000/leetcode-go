package next_greater_element_ii

// https://leetcode.cn/problems/next-greater-element-ii/

func nextGreaterElements(nums []int) []int {
	nums = append(nums, nums...)
	result := make([]int, len(nums))
	result[len(nums)-1] = -1

	// dp[i]表示nums[i]右边第一个比nums[i]更大的元素下标
	dp := make([]int, len(nums))
	dp[len(nums)-1] = -1
	for i := len(nums) - 2; i >= 0; i-- {
		idx := i + 1
		for idx != -1 && nums[idx] <= nums[i] {
			idx = dp[idx]
		}
		dp[i] = idx
		if dp[i] == -1 {
			result[i] = -1
		} else {
			result[i] = nums[dp[i]]
		}
	}

	return result[0 : len(result)/2]
}
