package longest_consecutive_sequence

import "math"

// https://leetcode.cn/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, n := range nums {
		set[n] = true
	}

	maxLen := 0
	for _, n := range nums {
		if _, exist := set[n-1]; exist {
			continue
		}

		cnt := 1
		x := n + 1
		for {
			if _, exist := set[x]; !exist {
				break
			}
			cnt++
			x++
		}

		maxLen = int(math.Max(float64(maxLen), float64(cnt)))
	}

	return maxLen
}
