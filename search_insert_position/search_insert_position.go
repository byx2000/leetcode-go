package search_insert_position

// https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	left, right, result := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			result = mid
			right = mid - 1
		} else {
			return mid
		}
	}
	return result
}
