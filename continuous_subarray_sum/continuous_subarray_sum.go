package continuous_subarray_sum

// https://leetcode.cn/problems/continuous-subarray-sum/

func checkSubarraySum(nums []int, k int) bool {
	// sum[i]表示nums[0~i-1]的和
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	set := map[int]bool{}
	for i := 2; i <= len(nums); i++ {
		set[sum[i-2]%k] = true
		if _, exist := set[sum[i]%k]; exist {
			return true
		}
	}

	return false
}
