package satisfiability_of_equality_equations

// https://leetcode.cn/problems/satisfiability-of-equality-equations/

func equationsPossible(equations []string) bool {
	// 为每一个字母分配一个唯一编号
	index := map[uint8]int{}
	n := 0
	for _, e := range equations {
		if _, exist := index[e[0]]; !exist {
			index[e[0]] = n
			n++
		}
		if _, exist := index[e[3]]; !exist {
			index[e[3]] = n
			n++
		}
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// 获取节点i所属集合
	var getParent func(i int) int
	getParent = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = getParent(parent[i])
		return parent[i]
	}

	// 合并i节点和j节点
	union := func(i, j int) {
		parent[getParent(i)] = getParent(j)
	}

	// 合并所有==连接的变量
	for _, e := range equations {
		if e[1:3] == "==" {
			union(index[e[0]], index[e[3]])
		}
	}

	// 检查所有!=连接的变量是否符合要求
	for _, e := range equations {
		if e[1:3] == "!=" {
			if getParent(index[e[0]]) == getParent(index[e[3]]) {
				return false
			}
		}
	}

	return true
}
