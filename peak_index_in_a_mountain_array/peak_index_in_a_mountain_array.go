package peak_index_in_a_mountain_array

// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	left, right, result := 1, len(arr)-2, -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid-1] {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
