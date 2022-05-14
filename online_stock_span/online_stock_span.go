package online_stock_span

// https://leetcode.cn/problems/online-stock-span/

type StockSpanner struct {
	prices []int
	dp     []int // dp[i]表示prices[i]左边第一个大于prices[i]的元素坐标
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (p *StockSpanner) Next(price int) int {
	if len(p.prices) == 0 {
		p.prices = append(p.prices, price)
		p.dp = append(p.dp, -1)
	} else {
		i := len(p.prices) - 1
		for i != -1 && p.prices[i] <= price {
			i = p.dp[i]
		}
		p.dp = append(p.dp, i)
		p.prices = append(p.prices, price)
	}

	return len(p.prices) - 1 - p.dp[len(p.dp)-1]
}
