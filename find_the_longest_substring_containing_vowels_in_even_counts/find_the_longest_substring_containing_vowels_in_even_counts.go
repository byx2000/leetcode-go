package find_the_longest_substring_containing_vowels_in_even_counts

import "math"

// https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/

func findTheLongestSubstring(s string) int {
	// 五个元音字母出现的奇偶性（1为奇，0为偶）
	cA, cE, cI, cO, cU := 0, 0, 0, 0, 0
	m := map[int]int{0: -1}
	maxLen := 0
	for i, c := range s {
		switch c {
		case 'a':
			cA = (cA + 1) % 2
		case 'e':
			cE = (cE + 1) % 2
		case 'i':
			cI = (cI + 1) % 2
		case 'o':
			cO = (cO + 1) % 2
		case 'u':
			cU = (cU + 1) % 2
		}

		key := (cA << 4) + (cE << 3) + (cI << 2) + (cO << 1) + cU
		if v, exist := m[key]; exist {
			maxLen = int(math.Max(float64(maxLen), float64(i-v)))
		} else {
			m[key] = i
		}
	}

	return maxLen
}
