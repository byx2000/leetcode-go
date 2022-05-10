package word_break

// https://leetcode.cn/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	words := map[string]bool{}
	for _, w := range wordDict {
		words[w] = true
	}

	cache := make([]bool, len(s))
	mark := make([]bool, len(s))

	// dp(index)表示s[0]~s[index]能否用wordDict中的单词表示
	var dp func(index int) bool
	dp = func(index int) (ret bool) {
		if mark[index] {
			return cache[index]
		}

		defer func() {
			mark[index] = true
			cache[index] = ret
		}()

		if _, exist := words[s[0:index+1]]; exist {
			return true
		}

		for i := 0; i <= index-1; i++ {
			if _, exist := words[s[i+1:index+1]]; exist && dp(i) {
				return true
			}
		}

		return false
	}

	return dp(len(s) - 1)
}
