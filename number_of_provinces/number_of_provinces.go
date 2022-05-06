package number_of_provinces

// https://leetcode-cn.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
	flag := make([]bool, len(isConnected))

	var dfs func(n int)
	dfs = func(n int) {
		if flag[n] {
			return
		}

		flag[n] = true
		for m := 0; m < len(isConnected[n]); m++ {
			if isConnected[n][m] == 1 {
				dfs(m)
			}
		}
	}

	cnt := 0
	for i := 0; i < len(isConnected); i++ {
		if !flag[i] {
			cnt++
			dfs(i)
		}
	}

	return cnt
}
