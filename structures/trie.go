package structures

import (
	"sort"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (t *TrieNode) InsertAll(words []string) {
	for _, word := range words {
		t.Insert(word)
	}
}

func (t *TrieNode) Insert(word string) {
	var (
		cur  = t
		next *TrieNode
		ok   bool
	)

	for _, c := range word {
		next, ok = cur.children[c]
		if !ok {
			next = NewTrie()
			cur.children[c] = next
		}

		cur = next
	}

	cur.isWord = true
}

func (t *TrieNode) Contains(word string) bool {
	var (
		cur  = t
		next *TrieNode
		ok   bool
	)

	for _, c := range word {
		next, ok = cur.children[c]
		if !ok {
			return false
		}

		cur = next
	}

	return cur.isWord
}

func (t *TrieNode) Search(searchWord string, limit int) []string {
	if limit < 1 {
		return nil
	}

	var (
		matches []string
		cur     = t
		next    *TrieNode
		ok      bool
	)

	for _, c := range searchWord {
		next, ok = cur.children[c]
		if !ok {
			return matches
		}

		cur = next
	}

	if cur.isWord {
		matches = append(matches, searchWord)
	}
	for c, child := range cur.children {
		matches = append(matches, child.getSuggestions(searchWord+string(c))...)
	}

	sort.Strings(matches)

	if len(matches) <= limit {
		return matches
	}

	return matches[:limit]
}

func (t *TrieNode) getSuggestions(prefix string) []string {
	var suggestions []string

	for c, child := range t.children {
		if child.isWord {
			suggestions = append(suggestions, prefix+string(c))
		}

		suggestions = append(suggestions, child.getSuggestions(prefix+string(c))...)
	}

	return suggestions
}

func (t *TrieNode) Remove(word string) bool {
	if len(word) == 0 {
		return false
	}

	var (
		cur            = t
		next           *TrieNode
		ok             bool
		deleteToken    = rune(word[0])
		depthDeleteIdx int
	)

	for i, c := range word {
		if (cur.isWord && !cur.isLeaf()) || len(cur.children) > 1 {
			depthDeleteIdx = i
			deleteToken = c
		}

		next, ok = cur.children[c]
		if !ok {
			return false
		}

		cur = next
	}

	if !cur.isLeaf() || !cur.isWord {
		return false
	}

	cur = t
	for i := 0; i < depthDeleteIdx; i++ {
		cur = cur.children[rune(word[i])]
	}

	delete(cur.children, deleteToken)
	return true
}

func (t *TrieNode) isLeaf() bool {
	return len(t.children) == 0
}

func (t *TrieNode) getNode(searchWord string) *TrieNode {
	var (
		cur  = t
		next *TrieNode
		ok   bool
	)

	for _, c := range searchWord {
		next, ok = cur.children[c]
		if !ok {
			return nil
		}

		cur = next
	}

	return cur
}
