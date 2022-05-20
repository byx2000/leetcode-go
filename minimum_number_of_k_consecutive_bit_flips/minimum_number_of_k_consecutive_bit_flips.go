package minimum_number_of_k_consecutive_bit_flips

// https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips/

func minKBitFlips(nums []int, k int) int {
	// 遍历nums，遇到0就翻转后面k个元素
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if i+k > len(nums) {
				return -1
			}
			cnt++
			for j := i; j < i+k; j++ {
				nums[j] = (nums[j] + 1) % 2
			}
		}
	}
	return cnt
}
