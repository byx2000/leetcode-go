package minimum_cost_to_cut_a_stick

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	cuts = append([]int{0}, cuts...)
	cuts = append(cuts, n)

	cache := make([][]int, len(cuts))
	mark := make([][]bool, len(cuts))
	for i := 0; i < len(cuts); i++ {
		cache[i] = make([]int, len(cuts))
		mark[i] = make([]bool, len(cuts))
	}

	// 切割cuts[i~j]范围内的棍子所需最小成本，不包括端点
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i+1 == j {
			return 0
		}

		result := math.MaxInt
		for k := i + 1; k < j; k++ {
			result = int(math.Min(float64(result), float64(dp(i, k)+dp(k, j)+cuts[j]-cuts[i])))
		}

		return result
	}

	return dp(0, len(cuts)-1)
}
