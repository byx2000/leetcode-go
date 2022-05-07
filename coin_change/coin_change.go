package coin_change

import (
	"math"
)

// https://leetcode-cn.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	cache := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		cache[i] = make([]int, amount+1)
	}

	// dp(index, sum)表示使用coins[0]~coins[index]面额的硬币凑够sum元的最小硬币个数
	var dp func(index, sum int) int
	dp = func(index, sum int) (ret int) {
		defer func() {
			cache[index][sum] = ret
		}()

		if cache[index][sum] != 0 {
			return cache[index][sum]
		}

		if index == 0 {
			if sum%coins[0] == 0 {
				return sum / coins[0]
			}
			return -1
		}

		if coins[index] > sum {
			return dp(index-1, sum)
		}

		r1 := dp(index-1, sum)
		r2 := dp(index, sum-coins[index])
		if r1 == -1 && r2 == -1 {
			return -1
		} else if r1 == -1 {
			return r2 + 1
		} else if r2 == -1 {
			return r1
		} else {
			return int(math.Min(float64(r1), float64(1+r2)))
		}
	}

	return dp(len(coins)-1, amount)
}
