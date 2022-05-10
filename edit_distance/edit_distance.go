package edit_distance

// https://leetcode.cn/problems/edit-distance/

func minDistance(s1 string, s2 string) int {
	min := func(a int, args ...int) int {
		r := a
		for _, n := range args {
			if n < r {
				r = n
			}
		}
		return r
	}

	cache := make([][]int, len(s1))
	mark := make([][]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		cache[i] = make([]int, len(s2))
		mark[i] = make([]bool, len(s2))
	}

	// dp(i, j)表示s1[0~i]和s2[0~j]的编辑距离
	var dp func(i, j int) int
	dp = func(i, j int) (ret int) {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if mark[i][j] {
			return cache[i][j]
		}

		defer func() {
			mark[i][j] = true
			cache[i][j] = ret
		}()

		if s1[i] == s2[j] {
			return dp(i-1, j-1)
		}

		return min(1+dp(i-1, j-1), 1+dp(i-1, j), 1+dp(i, j-1))
	}

	return dp(len(s1)-1, len(s2)-1)
}
