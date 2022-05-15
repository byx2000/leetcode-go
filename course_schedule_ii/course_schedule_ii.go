package course_schedule_ii

// https://leetcode.cn/problems/course-schedule-ii/

func findOrder(n int, edges [][]int) []int {
	// 构造邻接表
	// adj[i][j]表示j节点依赖i节点
	adj := map[int][]int{}
	for _, e := range edges {
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 计算出度
	out := make([]int, n)
	for _, v := range adj {
		for _, m := range v {
			out[m]++
		}
	}

	// 收集出度为0的节点
	var ready []int
	for i := 0; i < n; i++ {
		if out[i] == 0 {
			ready = append(ready, i)
		}
	}

	// 拓扑排序
	var order []int
	for len(ready) > 0 {
		cur := ready[0]
		order = append(order, cur)
		ready = ready[1:]
		for _, m := range adj[cur] {
			out[m]--
			if out[m] == 0 {
				ready = append(ready, m)
			}
		}
	}

	if len(order) == n {
		return order
	}
	return []int{}
}
