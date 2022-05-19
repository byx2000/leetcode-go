package subarray_sums_divisible_by_k

// https://leetcode.cn/problems/subarray-sums-divisible-by-k/

func subarraysDivByK(nums []int, k int) int {
	// a % b
	rem := func(a, b int) int {
		r := a % b
		if r < 0 {
			return (b + r) % b
		}
		return r
	}

	sum := 0
	m := map[int]int{0: 1}
	cnt := 0

	for _, n := range nums {
		sum += n
		cnt += m[rem(sum, k)]
		m[rem(sum, k)]++
	}

	return cnt
}
