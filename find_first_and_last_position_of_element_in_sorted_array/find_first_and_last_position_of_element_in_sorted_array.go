package find_first_and_last_position_of_element_in_sorted_array

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	findFirst := func() int {
		left, right, result := 0, len(nums)-1, -1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				result = mid
				right = mid - 1
			}
		}
		return result
	}

	findLast := func() int {
		left, right, result := 0, len(nums)-1, -1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				result = mid
				left = mid + 1
			}
		}
		return result
	}

	return []int{findFirst(), findLast()}
}
