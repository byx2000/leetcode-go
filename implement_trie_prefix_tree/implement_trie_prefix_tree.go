package implement_trie_prefix_tree

// https://leetcode.cn/problems/implement-trie-prefix-tree/

type node struct {
	m      [26]*node
	isWord bool
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{&node{}}
}

func (this *Trie) Insert(word string) {
	n := this.root
	for _, c := range word {
		if n.m[c-'a'] == nil {
			n.m[c-'a'] = &node{}
		}
		n = n.m[c-'a']
	}
	n.isWord = true
}

func (this *Trie) Search(word string) bool {
	n := this.root
	for _, c := range word {
		if n.m[c-'a'] != nil {
			n = n.m[c-'a']
		} else {
			return false
		}
	}
	return n.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	n := this.root
	for _, c := range prefix {
		if n.m[c-'a'] != nil {
			n = n.m[c-'a']
		} else {
			return false
		}
	}
	return true
}
