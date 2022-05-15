package course_schedule

// https://leetcode.cn/problems/course-schedule/

func canFinish(n int, edges [][]int) bool {
	// 构造邻接表
	adj := map[int][]int{}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}

	// 计算入度
	in := make([]int, n)
	for _, v := range adj {
		for _, m := range v {
			in[m]++
		}
	}

	// 收集入度为0的节点
	var ready []int
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			ready = append(ready, i)
		}
	}

	//拓扑排序
	cnt := 0
	for len(ready) > 0 {
		cnt++
		cur := ready[0]
		ready = ready[1:]
		for _, m := range adj[cur] {
			in[m]--
			if in[m] == 0 {
				ready = append(ready, m)
			}
		}
	}

	return cnt == n
}
