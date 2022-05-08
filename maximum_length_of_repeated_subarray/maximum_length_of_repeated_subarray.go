package maximum_length_of_repeated_subarray

import "math"

// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

func findLength(nums1 []int, nums2 []int) int {
	cache := make([][]int, len(nums1))
	mark := make([][]bool, len(nums1))
	for i := 0; i < len(nums1); i++ {
		cache[i] = make([]int, len(nums2))
		mark[i] = make([]bool, len(nums2))
	}

	// dp(i, j)表示以nums1[i]和nums2[j]结尾的最长公共子数组长度
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i == 0 || j == 0 {
			if nums1[i] == nums2[j] {
				return 1
			}
			return 0
		}
		if nums1[i] == nums2[j] {
			return 1 + dp(i-1, j-1)
		}
		return 0
	}

	maxLen := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			maxLen = int(math.Max(float64(maxLen), float64(dp(i, j))))
		}
	}

	return maxLen
}
