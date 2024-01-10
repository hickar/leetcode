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

func TestTrieIsWord(t *testing.T) {
	trie := NewTrie()
	trie.InsertAll([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"})

	assert.NotPanics(t, func() {
		assert.True(t, trie.getNode("mouse").isWord)
		assert.True(t, trie.getNode("mousepad").isWord)
	})
}

func TestTrieSearch(t *testing.T) {
	input := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	trie := NewTrie()
	trie.InsertAll(input)

	expected := []string{"mouse", "mousepad"}
	output := trie.Search("mouse", 3)

	assert.Equal(t, expected, output)
}

func TestTrieContains(t *testing.T) {
	trie := NewTrie()
	trie.Insert("trie")
	trie.Insert("tree")

	assert.True(t, trie.Contains("trie"))
	assert.True(t, trie.Contains("tree"))
	assert.False(t, trie.Contains("tr"))
}

func TestTrieRemove(t *testing.T) {
	trie := NewTrie()
	trie.Insert("trie")
	trie.Insert("tree")

	assert.False(t, trie.Remove("tr"))

	assert.True(t, trie.Remove("tree"))
	assert.False(t, trie.Contains("tree"))

	assert.True(t, trie.Contains("trie"))
	assert.True(t, trie.Remove("trie"))
	assert.False(t, trie.Contains("trie"))

	assert.Len(t, trie.children, 0)
}
