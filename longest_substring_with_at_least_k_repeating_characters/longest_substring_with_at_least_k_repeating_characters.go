package longest_substring_with_at_least_k_repeating_characters

import "math"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/

func longestSubstring(s string, k int) int {
	// 判断每个字符出现的个数是否都不少于k
	var ok = func(chs map[int]int) bool {
		for _, v := range chs {
			if v < k {
				return false
			}
		}
		return true
	}

	maxLen := 0
	for i := 0; i < len(s); i++ {
		chs := map[int]int{}
		for j := i; j < len(s); j++ {
			v, exist := chs[int(s[j])]
			if exist {
				chs[int(s[j])] = v + 1
			} else {
				chs[int(s[j])] = 1
			}
			if ok(chs) {
				maxLen = int(math.Max(float64(maxLen), float64(j-i+1)))
			}
		}
	}

	return maxLen
}
