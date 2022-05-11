package longest_valid_parentheses

import "math"

// https://leetcode.cn/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	maxLen := 0

	for i := 0; i < len(s); i++ {
		// 剪枝
		if s[i] == ')' || len(s)-i-1 <= maxLen {
			continue
		}

		// 计算以s[i]开头的最长有效括号序列长度
		var stack []int
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				stack = append(stack, '(')
			} else {
				if len(stack) > 0 && stack[len(stack)-1] == '(' {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = append(stack, ')')
				}
			}
			if len(stack) == 0 {
				maxLen = int(math.Max(float64(maxLen), float64(j-i+1)))
			}
		}
	}

	return maxLen
}
