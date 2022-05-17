package design_add_and_search_words_data_structure

// https://leetcode.cn/problems/design-add-and-search-words-data-structure/

type node struct {
	m      [26]*node
	isWord bool
}

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{&node{}}
}

func (this *WordDictionary) AddWord(word string) {
	n := this.root
	for _, c := range word {
		if n.m[c-'a'] == nil {
			n.m[c-'a'] = &node{}
		}
		n = n.m[c-'a']
	}
	n.isWord = true
}

func search(word string, index int, n *node) bool {
	if index == len(word) {
		return n.isWord
	}

	c := word[index]
	if c == '.' {
		for _, nn := range n.m {
			if nn != nil && search(word, index+1, nn) {
				return true
			}
		}
		return false
	} else {
		if n.m[c-'a'] != nil {
			return search(word, index+1, n.m[c-'a'])
		} else {
			return false
		}
	}
}

func (this *WordDictionary) Search(word string) bool {
	return search(word, 0, this.root)
}
