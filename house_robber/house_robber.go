package house_robber

import "math"

// https://leetcode.cn/problems/house-robber/

func rob(nums []int) int {
	cache := make([]int, len(nums))
	mark := make([]bool, len(nums))

	// dp(i)表示打劫nums[0~i]所能得到的最大金额
	var dp func(i int) int
	dp = func(i int) (ret int) {
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

		return int(math.Max(float64(dp(i-1)), float64(nums[i]+dp(i-2))))
	}

	return dp(len(nums) - 1)
}
