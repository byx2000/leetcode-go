package word_ladder

// https://leetcode.cn/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 如果endWord不在词典中，则直接返回0
	findEnd := false
	for _, w := range wordList {
		if w == endWord {
			findEnd = true
			break
		}
	}
	if !findEnd {
		return 0
	}

	// 判断s1和s2是否仅相差一个字符
	match := func(s1, s2 string) bool {
		if len(s1) != len(s2) {
			return false
		}

		flag := false
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if flag {
					return false
				}
				flag = true
			}
		}

		return true
	}

	// 判断wordList[i]和wordList[j]是否仅相差一个字符
	cache := make([][]bool, len(wordList))
	mark := make([][]bool, len(wordList))
	for i := 0; i < len(wordList); i++ {
		cache[i] = make([]bool, len(wordList))
		mark[i] = make([]bool, len(wordList))
	}
	matchIndex := func(i, j int) bool {
		if mark[i][j] {
			return cache[i][j]
		}
		mark[i][j] = true
		cache[i][j] = match(wordList[i], wordList[j])
		return cache[i][j]
	}

	// 初始化队列
	var queue []int
	visited := map[int]bool{}
	for i, w := range wordList {
		if match(beginWord, w) {
			queue = append(queue, i)
			visited[i] = true
		}
	}

	// BFS
	step := 2
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			cur := queue[0]
			queue = queue[1:]

			if wordList[cur] == endWord {
				return step
			}

			for j := range wordList {
				if !visited[j] && matchIndex(cur, j) {
					queue = append(queue, j)
					visited[j] = true
				}
			}
		}
		step++
	}

	return 0
}
