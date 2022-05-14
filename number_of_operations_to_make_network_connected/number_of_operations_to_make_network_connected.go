package number_of_operations_to_make_network_connected

// https://leetcode.cn/problems/number-of-operations-to-make-network-connected/

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	// 获取节点i所属集合
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// 合并i节点和j节点
	var getParent func(i int) int
	getParent = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = getParent(parent[i])
		return parent[i]
	}

	union := func(i, j int) {
		parent[getParent(i)] = getParent(j)
	}

	for _, c := range connections {
		union(c[0], c[1])
	}

	// 统计连通分量的数量
	set := map[int]bool{}
	for i := 0; i < n; i++ {
		set[getParent(i)] = true
	}

	return len(set) - 1
}
