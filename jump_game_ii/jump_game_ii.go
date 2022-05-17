package jump_game_ii

import "math"

// https://leetcode.cn/problems/jump-game-ii/

func jump(nums []int) int {
	cache := make([]int, len(nums))
	mark := make([]bool, len(nums))

	// dp(index)表示从nums[index]到最后一个位置的最小跳跃数
	var dp func(index int) int
	dp = func(index int) (ret int) {
		if mark[index] {
			return cache[index]
		}

		defer func() {
			mark[index] = true
			cache[index] = ret
		}()

		if index == len(nums)-1 {
			return 0
		}

		if nums[index] >= len(nums)-index-1 {
			return 1
		}

		result := math.MaxInt
		for i := 1; i <= nums[index]; i++ {
			r := dp(index + i)
			if r != math.MaxInt {
				result = int(math.Min(float64(result), float64(r+1)))
			}
		}

		return result
	}

	return dp(0)
}
