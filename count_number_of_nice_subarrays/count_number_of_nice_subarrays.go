package count_number_of_nice_subarrays

// https://leetcode.cn/problems/count-number-of-nice-subarrays/

func numberOfSubarrays(nums []int, k int) int {
	// 奇数变1，偶数变0
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] % 2
	}

	// 求区间和为k的子数组数量

	sum := 0
	m := map[int]int{0: 1}
	cnt := 0

	for _, n := range nums {
		sum += n
		cnt += m[sum-k]
		m[sum]++
	}

	return cnt
}
