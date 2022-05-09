package minimum_size_subarray_sum

import (
	"math"
)

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

func minSubArrayLen1(target int, nums []int) int {
	i, j, sum := 0, 0, 0
	minLen := math.MaxInt

	for j < len(nums) {
		if sum >= target {
			minLen = int(math.Min(float64(minLen), float64(j-i)))
			sum -= nums[i]
			i++
		} else {
			sum += nums[j]
			j++
		}
	}

	for sum >= target {
		minLen = int(math.Min(float64(minLen), float64(j-i)))
		sum -= nums[i]
		i++
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func minSubArrayLen2(target int, nums []int) int {
	i, j, sum := 0, 0, 0
	minLen := math.MaxInt

	// [i, j)区间和 < target
	for j < len(nums) {
		// 向右移动j，直到sum >= target
		for j < len(nums) && sum < target {
			sum += nums[j]
			j++
		}

		//向右移动i，直到sum < target
		for sum >= target {
			minLen = int(math.Min(float64(minLen), float64(j-i)))
			sum -= nums[i]
			i++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
