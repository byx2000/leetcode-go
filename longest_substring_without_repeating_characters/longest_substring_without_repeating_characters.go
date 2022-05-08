package longest_substring_without_repeating_characters

import "math"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	i, j, maxLen := 0, 0, 0
	set := map[uint8]bool{}
	for j < len(s) {
		if _, exist := set[s[j]]; exist {
			maxLen = int(math.Max(float64(maxLen), float64(j-i)))
			delete(set, s[i])
			i++
		} else {
			set[s[j]] = true
			j++
		}
	}
	maxLen = int(math.Max(float64(maxLen), float64(j-i)))
	return maxLen
}
