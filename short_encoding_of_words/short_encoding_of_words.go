package short_encoding_of_words

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/short-encoding-of-words/

type node struct {
	m [26]*node
}

type trie struct {
	root *node
}

func (t trie) insert(word string) {
	n := t.root
	for _, c := range word {
		if n.m[c-'a'] == nil {
			n.m[c-'a'] = &node{}
		}
		n = n.m[c-'a']
	}
}

func (t trie) search(word string) bool {
	n := t.root
	for _, c := range word {
		if n.m[c-'a'] != nil {
			n = n.m[c-'a']
		} else {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

func minimumLengthEncoding(words []string) int {
	// 将所有单词按照长度倒序排序
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[j]) < len(words[i])
	})
	fmt.Println(words)

	// 构建字典树
	t := trie{&node{}}
	result := 0
	for _, w := range words {
		rw := reverse(w)
		if !t.search(rw) {
			t.insert(rw)
			result += len(w) + 1
		}
	}

	return result
}
