package wildcard_matching

// https://leetcode.cn/problems/wildcard-matching/

func isMatch(s string, p string) (ret bool) {
	cache := make([][]bool, len(s))
	mark := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]bool, len(p))
		mark[i] = make([]bool, len(p))
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) (ret bool) {
		if i < len(s) && j < len(p) && mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			if i < len(s) && j < len(p) {
				mark[i][j] = true
				cache[i][j] = ret
			}
		}()

		if j == len(p) {
			return i == len(s)
		}

		if p[j] == '?' {
			return dfs(i+1, j+1)
		} else if p[j] == '*' {
			for k := i; k <= len(s); k++ {
				if dfs(k, j+1) {
					return true
				}
			}
			return false
		} else {
			if i < len(s) && s[i] == p[j] {
				return dfs(i+1, j+1)
			}
			return false
		}
	}

	return dfs(0, 0)
}
