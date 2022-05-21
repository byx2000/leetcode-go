package minimum_window_substring

import "math"

// https://leetcode.cn/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	m1 := map[uint8]int{}
	for _, c := range t {
		m1[uint8(c)]++
	}

	m2 := map[uint8]int{}
	match := func() bool {
		for k, v := range m1 {
			if m2[k] < v {
				return false
			}
		}
		return true
	}

	i, j := 0, 0
	minLen := math.MaxInt
	start, end := 0, 0

	for j < len(s) {
		for j < len(s) && !match() {
			m2[s[j]]++
			j++
		}

		for match() {
			if j-i < minLen {
				minLen = j - i
				start, end = i, j
			}
			m2[s[i]]--
			i++
		}
	}

	return s[start:end]
}
