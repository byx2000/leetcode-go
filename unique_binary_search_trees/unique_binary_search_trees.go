package unique_binary_search_trees

// https://leetcode.cn/problems/unique-binary-search-trees/

func numTrees(n int) int {
	cache := make([]int, n+1)
	mark := make([]bool, n+1)

	var dp func(n int) int
	dp = func(n int) (ret int) {
		if mark[n] {
			return cache[n]
		}

		defer func() {
			mark[n] = true
			cache[n] = ret
		}()

		if n == 1 {
			return 1
		}

		// i作为根节点，[1, i-1]为左子树，[i+1, n]为右子树
		cnt := 0
		for i := 2; i <= n-1; i++ {
			cnt += dp(i-1) * dp(n-i)
		}

		return cnt + 2*dp(n-1)
	}

	return dp(n)
}
