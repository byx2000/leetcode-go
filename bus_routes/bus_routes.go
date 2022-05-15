package bus_routes

// https://leetcode.cn/problems/bus-routes/

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// 构造邻接表
	adj := make([][]bool, len(routes))
	for i := 0; i < len(routes); i++ {
		adj[i] = make([]bool, len(routes))
	}
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes); j++ {
			if i != j {
				set := map[int]bool{}
				for _, p := range routes[i] {
					set[p] = true
				}
				for _, p := range routes[j] {
					if _, exist := set[p]; exist {
						adj[i][j] = true
						adj[j][i] = true
						break
					}
				}
			}
		}
	}

	var queue []int
	visited := map[int]bool{}

	// 构造起点集合
	for i, r := range routes {
		for _, p := range r {
			if p == source {
				queue = append(queue, i)
				visited[i] = true
				break
			}
		}
	}

	// 构造终点集合
	targetSet := map[int]bool{}
	for i, r := range routes {
		for _, p := range r {
			if p == target {
				targetSet[i] = true
				break
			}
		}
	}

	// BFS
	step := 1
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			cur := queue[0]
			queue = queue[1:]
			if _, exist := targetSet[cur]; exist {
				return step
			}

			for next, v := range adj[cur] {
				if v {
					if _, exist := visited[next]; !exist {
						queue = append(queue, next)
						visited[next] = true
					}
				}
			}
		}
		step++
	}

	return -1
}

func numBusesToDestination2(routes [][]int, source int, target int) int {
	// 记录每个站台对应的公交线路编号
	lines := map[int][]int{}
	for i, r := range routes {
		for _, p := range r {
			lines[p] = append(lines[p], i)
		}
	}

	// 添加初始线路集合
	queue := []int{source}
	visited := map[int]bool{source: true}

	// BFS
	step := 0
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}

			for _, r := range lines[cur] {
				for _, p := range routes[r] {
					if _, exist := visited[p]; !exist {
						queue = append(queue, p)
						visited[p] = true
					}
				}
			}
		}
		step++
	}

	return -1
}
