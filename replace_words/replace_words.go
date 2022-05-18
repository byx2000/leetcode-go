package replace_words

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/replace-words/

func replaceWords(dict []string, sentence string) string {
	// 将所有词根按照长度排序
	sort.SliceStable(dict, func(i, j int) bool {
		return len(dict[i]) < len(dict[j])
	})

	words := strings.Split(sentence, " ")
	var roots []string
	for _, w := range words {
		find := false
		for _, r := range dict {
			if strings.HasPrefix(w, r) {
				find = true
				roots = append(roots, r)
				break
			}
		}
		if !find {
			roots = append(roots, w)
		}
	}

	return strings.Join(roots, " ")
}
