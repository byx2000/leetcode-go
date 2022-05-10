package partition_array_for_maximum_sum

import (
	"math"
)

// https://leetcode.cn/problems/partition-array-for-maximum-sum/

func maxSumAfterPartitioning(arr []int, k int) int {
	max := func(i, j int) int {
		result := arr[i]
		for k := i + 1; k <= j; k++ {
			result = int(math.Max(float64(result), float64(arr[k])))
		}
		return result
	}

	cache := make([]int, len(arr))
	mark := make([]bool, len(arr))

	// dp(index)表示分隔arr[0]~arr[index]
	var dp func(index int) int
	dp = func(index int) (ret int) {
		if mark[index] {
			return cache[index]
		}

		defer func() {
			mark[index] = true
			cache[index] = ret
		}()

		if index == 0 {
			return arr[0]
		}

		// 将arr[0~index]分隔成arr[0~i]和arr[i+1~index]
		maxVal := arr[index]
		result := math.MinInt
		for i := index - 1; index-i <= k && i >= 0; i-- {
			result = int(math.Max(float64(result), float64(maxVal*(index-i)+dp(i))))
			maxVal = int(math.Max(float64(maxVal), float64(arr[i])))
		}

		// arr[0~index]作为一组
		if index+1 <= k {
			result = int(math.Max(float64(result), float64(max(0, index)*(index+1))))
		}

		return result
	}

	return dp(len(arr) - 1)
}
