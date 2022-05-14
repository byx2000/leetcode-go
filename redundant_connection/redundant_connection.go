package redundant_connection

// https://leetcode.cn/problems/redundant-connection/

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)

	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	// 获取节点x所属集合
	var getParent func(x int) int
	getParent = func(x int) int {
		if parent[x] == x {
			return x
		}
		parent[x] = getParent(parent[x])
		return parent[x]
	}

	// 合并x节点和y节点
	union := func(x, y int) {
		parent[getParent(x)] = getParent(y)
	}

	for _, e := range edges {
		if getParent(e[0]) == getParent(e[1]) {
			return e
		}
		union(e[0], e[1])
	}

	return nil
}
