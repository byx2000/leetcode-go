package all_paths_from_source_to_target

// https://leetcode-cn.com/problems/all-paths-from-source-to-target/

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	mem := make([]bool, n)
	var paths [][]int

	var dfs func(cur int, path []int)
	dfs = func(cur int, path []int) {
		if cur == n-1 {
			path = append(path, cur)
			p := make([]int, len(path))
			copy(p, path)
			paths = append(paths, p)
			return
		}

		if mem[cur] {
			return
		}

		mem[cur] = true
		for _, next := range graph[cur] {
			dfs(next, append(path, cur))
		}
		mem[cur] = false
	}

	dfs(0, []int{})

	return paths
}
