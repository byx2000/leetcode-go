package route_between_nodes_lcci

// https://leetcode-cn.com/problems/route-between-nodes-lcci/

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 构造邻接表
	adj := map[int][]int{}
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, e := range graph {
		adj[e[0]] = append(adj[e[0]], e[1])
	}

	mark := make([]bool, n)
	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		if cur == target {
			return true
		}

		if mark[cur] {
			return false
		}

		mark[cur] = true
		defer func() {
			mark[cur] = false
		}()

		for _, next := range adj[cur] {
			if dfs(next) {
				return true
			}
		}

		return false
	}

	return dfs(start)
}
