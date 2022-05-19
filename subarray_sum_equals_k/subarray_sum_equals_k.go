package subarray_sum_equals_k

// https://leetcode.cn/problems/subarray-sum-equals-k/

func subarraySum1(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

func subarraySum2(nums []int, k int) int {
	sum := 0
	m := map[int]int{0: 1}
	cnt := 0

	// sum表示nums[0~i]的和
	// m[s]记录了和为s的nums[0~j]子数组个数（j < i）
	for _, n := range nums {
		sum += n
		cnt += m[sum-k]
		m[sum]++
	}

	return cnt
}
