package palindromic_substrings

// https://leetcode.cn/problems/palindromic-substrings/

func countSubstrings(s string) int {
	cache := make([][]bool, len(s))
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]bool, len(s))
		mark[i] = make([]bool, len(s))
	}

	// 判断s[i~j]是不是回文串
	var isPalindromic func(i, j int) bool
	isPalindromic = func(i, j int) (ret bool) {
		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if i == j {
			return true
		}

		if i+1 == j {
			return s[i] == s[j]
		}

		return s[i] == s[j] && isPalindromic(i+1, j-1)
	}

	cnt := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindromic(i, j) {
				cnt++
			}
		}
	}

	return cnt
}
