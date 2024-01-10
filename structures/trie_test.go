package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("trie")
	trie.Insert("tree")

	assert.Contains(t, trie.children, 't')
	assert.Len(t, trie.children, 1)
	assert.Contains(t, trie.children['t'].children, 'r')
	assert.Len(t, trie.children['t'].children, 1)

	assert.Contains(t, trie.children['t'].children['r'].children, 'i')
	assert.Contains(t, trie.children['t'].children['r'].children, 'e')
	assert.Len(t, trie.children['t'].children['r'].children, 2)
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()
	trie.Insert("trie")
	trie.Insert("tree")

	assert.True(t, trie.Search("trie"))
	assert.True(t, trie.Search("tree"))
	assert.False(t, trie.Search("tr"))
}

func TestTrieRemove(t *testing.T) {
	trie := NewTrie()
	trie.Insert("trie")
	trie.Insert("tree")

	assert.False(t, trie.Remove("tr"))

	assert.True(t, trie.Remove("tree"))
	assert.False(t, trie.Search("tree"))

	assert.True(t, trie.Search("trie"))
	assert.True(t, trie.Remove("trie"))
	assert.False(t, trie.Search("trie"))

	assert.Len(t, trie.children, 0)
}
