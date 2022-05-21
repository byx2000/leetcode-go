package shortest_supersequence_lcci

import "math"

// https://leetcode.cn/problems/shortest-supersequence-lcci/

func shortestSeq(big []int, small []int) []int {
	set := map[int]bool{}
	for _, c := range small {
		set[c] = true
	}

	i, j := 0, 0
	s := map[int]int{}
	minLen := math.MaxInt
	start, end := -1, -1

	match := func() bool {
		for k := range set {
			if s[k] == 0 {
				return false
			}
		}
		return true
	}

	for j < len(big) {
		for j < len(big) && !match() {
			s[big[j]]++
			j++
		}

		for match() {
			if j-i < minLen {
				minLen = j - i
				start, end = i, j-1
			}
			s[big[i]]--
			i++
		}
	}

	if start < 0 {
		return []int{}
	}
	return []int{start, end}
}
