package structures

type Trie struct {
	Nodes map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		Nodes: make(map[rune]*Trie),
	}
}

func NewTrieWithWords(words []string) *Trie {
	trie := NewTrie()
	trie.AddWords(words)

	return trie
}

func (t *Trie) AddWords(words []string) {
	for _, word := range words {
		t.AddWord(word)
	}
}

func (t *Trie) AddWord(word string) {
	curNode := t

	for _, c := range word {
		nextNode, ok := curNode.Nodes[c]
		if !ok {
			nextNode = NewTrie()
			curNode.Nodes[c] = nextNode
		}

		curNode = nextNode
	}
}
