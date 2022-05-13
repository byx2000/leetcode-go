package largest_rectangle_in_histogram

import (
	"math"
)

// https://leetcode.cn/problems/largest-rectangle-in-histogram/

func largestRectangleArea(nums []int) int {
	// leftIdx[i]表示nums[i]左边第一个小于nums[i]的元素索引
	leftIdx := make([]int, len(nums))
	leftIdx[0] = -1
	for i := 1; i < len(nums); i++ {
		idx := i - 1
		for idx != -1 && nums[idx] >= nums[i] {
			idx = leftIdx[idx]
		}
		leftIdx[i] = idx
	}

	// rightIdx[i]表示nums[i]右边第一个小于nums[i]的元素索引
	rightIdx := make([]int, len(nums))
	rightIdx[len(nums)-1] = len(nums)
	for i := len(nums) - 2; i >= 0; i-- {
		idx := i + 1
		for idx != len(nums) && nums[idx] >= nums[i] {
			idx = rightIdx[idx]
		}
		rightIdx[i] = idx
	}

	maxArea := 0
	for i := 0; i < len(nums); i++ {
		maxArea = int(math.Max(float64(maxArea), float64(nums[i]*(rightIdx[i]-leftIdx[i]-1))))
	}

	return maxArea
}
