package coin_change_2

// https://leetcode-cn.com/problems/coin-change-2/

func change(amount int, coins []int) int {
	cache := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		cache[i] = make([]int, amount+1)
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = -1
		}
	}

	// dp(index, sum)表示coins[0]~coins[index]凑够sum元的组合数
	var dp func(index, sum int) int
	dp = func(index, sum int) (ret int) {
		defer func() {
			cache[index][sum] = ret
		}()

		if cache[index][sum] >= 0 {
			return cache[index][sum]
		}

		if index == 0 {
			if sum%coins[0] == 0 {
				return 1
			}
			return 0
		}

		if coins[index] > sum {
			return dp(index-1, sum)
		}

		return dp(index-1, sum) + dp(index, sum-coins[index])
	}

	return dp(len(coins)-1, amount)
}
