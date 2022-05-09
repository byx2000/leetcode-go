package longest_substring_without_repeating_characters

import "math"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring1(s string) int {
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

func lengthOfLongestSubstring2(s string) int {
	i, j, maxLen := 0, 0, 0
	set := map[uint8]bool{}

	// [i, j)区间内无重复字符
	for j < len(s) {
		// j向右移动，直到[i, j)有重复字符
		for j < len(s) {
			maxLen = int(math.Max(float64(maxLen), float64(j-i)))
			if _, exist := set[s[j]]; exist {
				break
			}

			set[s[j]] = true
			j++
		}

		if j == len(s) {
			break
		}

		// i向右移动，直到[i, j)没有重复字符
		for {
			if _, exist := set[s[j]]; !exist {
				break
			}
			delete(set, s[i])
			i++
		}
	}

	maxLen = int(math.Max(float64(maxLen), float64(j-i)))

	return maxLen
}
