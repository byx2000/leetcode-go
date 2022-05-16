package best_time_to_buy_and_sell_stock_ii

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit1(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}

func maxProfit2(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// in表示当前手里持有股票的最大获利
	// out表示当前手里未持有股票的最大获利
	in, out := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		in = max(in, out-prices[i])
		out = max(out, in+prices[i])
	}
	return max(in, out)
}
