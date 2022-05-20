package palindrome_partitioning

// https://leetcode.cn/problems/palindrome-partitioning/

func partition(s string) [][]string {
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

	var result [][]string

	var dfs func(index int, r []string)
	dfs = func(index int, r []string) {
		if index == len(s) {
			r1 := make([]string, len(r))
			copy(r1, r)
			result = append(result, r1)
			return
		}
		for i := index; i < len(s); i++ {
			if isPalindromic(index, i) {
				dfs(i+1, append(r, s[index:i+1]))
			}
		}
	}

	dfs(0, []string{})

	return result
}
