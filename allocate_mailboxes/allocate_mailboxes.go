package allocate_mailboxes

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/allocate-mailboxes/

func minDistance(houses []int, k int) int {
	// 将houses按从小到大的顺序排序
	sort.Ints(houses)

	// 求houses[i]~houses[j]共用一个油桶的最小距离
	minDis := func(i, j int) int {
		m := (i + j) / 2
		dis := 0
		for k := i; k < m; k++ {
			dis += houses[m] - houses[k]
		}
		for k := m + 1; k <= j; k++ {
			dis += houses[k] - houses[m]
		}
		return dis
	}

	cache := make([][]int, len(houses))
	mark := make([][]bool, len(houses))
	for i := 0; i < len(houses); i++ {
		cache[i] = make([]int, k+1)
		mark[i] = make([]bool, k+1)
	}

	// dp(index, k)表示将houses[0~index]分成k份
	// index+1 >= k
	var dp func(index, k int) int
	dp = func(index, k int) (ret int) {
		if mark[index][k] {
			return cache[index][k]
		}

		defer func() {
			mark[index][k] = true
			cache[index][k] = ret
		}()

		if k == 1 {
			return minDis(0, index)
		}

		// 将houses[0~index]拆分成houses[0~i]和houses[i+1~index]
		result := math.MaxInt
		for i := index - 1; i+1 >= k-1; i-- {
			result = int(math.Min(float64(result), float64(dp(i, k-1)+minDis(i+1, index))))
		}

		return result
	}

	return dp(len(houses)-1, k)
}
