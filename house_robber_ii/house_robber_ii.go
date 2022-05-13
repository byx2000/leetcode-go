package house_robber_ii

import "math"

// https://leetcode.cn/problems/house-robber-ii/

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	calc := func(nums []int, i int) int {
		cache := make([]int, len(nums))
		mark := make([]bool, len(nums))

		// dp(i)表示打劫nums[0~i]所能得到的最大金额
		var dp func(nums []int, i int) int
		dp = func(nums []int, i int) (ret int) {
			if mark[i] {
				return cache[i]
			}

			defer func() {
				mark[i] = true
				cache[i] = ret
			}()

			if i == 0 {
				return nums[0]
			}

			if i == 1 {
				return int(math.Max(float64(nums[0]), float64(nums[1])))
			}

			return int(math.Max(float64(dp(nums, i-1)), float64(nums[i]+dp(nums, i-2))))
		}

		return dp(nums, i)
	}

	return int(math.Max(float64(calc(nums[:len(nums)-1], len(nums)-2)), float64(calc(nums[1:], len(nums)-2))))
}
