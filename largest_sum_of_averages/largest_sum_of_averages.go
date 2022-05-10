package largest_sum_of_averages

import "math"

// https://leetcode.cn/problems/largest-sum-of-averages/

func largestSumOfAverages(nums []int, k int) float64 {
	cache := make([][]float64, len(nums))
	mark := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]float64, k+1)
		mark[i] = make([]bool, k+1)
	}

	// 将nums[0~index]分成k份
	var dp func(index, k int) float64
	dp = func(index, k int) (ret float64) {
		if mark[index][k] {
			return cache[index][k]
		}

		defer func() {
			mark[index][k] = true
			cache[index][k] = ret
		}()

		if index+1 == k {
			sum := 0
			for i := 0; i <= index; i++ {
				sum += nums[i]
			}
			return float64(sum)
		}

		if k == 1 {
			sum := 0
			for i := 0; i <= index; i++ {
				sum += nums[i]
			}
			return float64(sum) / float64(index+1)
		}

		// 将nums[0~index]拆分成nums[0~i]和nums[i+1~index]
		sum := nums[index]
		result := -1.0
		for i := index - 1; i+1 >= k-1 && i >= 0; i-- {
			result = math.Max(result, float64(sum)/float64(index-i)+dp(i, k-1))
			sum += nums[i]
		}

		return result
	}

	return dp(len(nums)-1, k)
}
