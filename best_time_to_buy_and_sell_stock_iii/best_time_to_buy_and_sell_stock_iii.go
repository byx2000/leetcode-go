package best_time_to_buy_and_sell_stock_iii

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/

func maxProfit(prices []int) int {
	max := func(a int, args ...int) int {
		r := a
		for _, n := range args {
			r = int(math.Max(float64(r), float64(n)))
		}
		return r
	}

	in1, out1, in2, out2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		in1 = max(in1, -prices[i])
		out1 = max(out1, in1+prices[i])
		in2 = max(in2, out1-prices[i])
		out2 = max(out2, in2+prices[i])
	}

	return max(0, out1, out2)
}
