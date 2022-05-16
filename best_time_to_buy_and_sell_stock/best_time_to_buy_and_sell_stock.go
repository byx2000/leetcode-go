package best_time_to_buy_and_sell_stock

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	// min[i]表示prices[0~i]的最小值
	min := make([]int, len(prices))
	min[0] = prices[0]

	result := 0
	for i := 1; i < len(prices); i++ {
		min[i] = int(math.Min(float64(min[i-1]), float64(prices[i])))
		result = int(math.Max(float64(result), float64(prices[i]-min[i])))
	}

	return result
}
