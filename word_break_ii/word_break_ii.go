package word_break_ii

import "strings"

// https://leetcode.cn/problems/word-break-ii/

func wordBreak(s string, wordDict []string) []string {
	words := map[string]bool{}
	for _, w := range wordDict {
		words[w] = true
	}

	var result []string

	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index >= len(s) {
			result = append(result, strings.Join(path, " "))
			return
		}

		for i := index; i < len(s); i++ {
			if _, exist := words[s[index:i+1]]; exist {
				dfs(i+1, append(path, s[index:i+1]))
			}
		}
	}

	dfs(0, []string{})

	return result
}
