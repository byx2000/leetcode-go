package length_of_longest_fibonacci_subsequence

import "math"

// https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/

func lenLongestFibSubseq(nums []int) int {
	cache := make([][]int, len(nums))
	mark := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]int, len(nums))
		mark[i] = make([]bool, len(nums))
	}

	// 存储nums元素值对应的索引
	m := map[int]int{}
	for i, v := range nums {
		m[v] = i
	}

	// dp(i, j)表示以nums[i]、nums[j]结尾的最长斐波拉契子序列长度
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		k, exist := m[nums[j]-nums[i]]
		if exist && k < i {
			return dp(k, i) + 1
		}

		return 2
	}

	maxLen := 2
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			maxLen = int(math.Max(float64(maxLen), float64(dp(i, j))))
		}
	}

	if maxLen == 2 {
		return 0
	}
	return maxLen
}
