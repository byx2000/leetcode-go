package design_add_and_search_words_data_structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.cn/problems/design-add-and-search-words-data-structure/

func Test1(t *testing.T) {
	d := Constructor()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")
	assert.False(t, d.Search("pad"))
	assert.True(t, d.Search("bad"))
	assert.True(t, d.Search(".ad"))
	assert.True(t, d.Search("b.."))
	assert.False(t, d.Search("b."))
	assert.False(t, d.Search("d..."))
	assert.False(t, d.Search(".d"))
	assert.False(t, d.Search("...d"))
}

func Test2(t *testing.T) {
	d := Constructor()
	d.AddWord("at")
	d.AddWord("and")
	d.AddWord("an")
	d.AddWord("add")
	assert.False(t, d.Search("a"))
	assert.False(t, d.Search(".at"))
	d.AddWord("bat")
	assert.True(t, d.Search(".at"))
	assert.True(t, d.Search("an."))
	assert.False(t, d.Search("a.d."))
}
