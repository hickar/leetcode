package structures

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
		curNode  = t
		nextNode *TrieNode
		ok       bool
	)

	for _, c := range word {
		nextNode, ok = curNode.children[c]
		if !ok {
			nextNode = NewTrie()
			curNode.children[c] = nextNode
		}

		curNode = nextNode
	}

	curNode.isWord = true
}

func (t *TrieNode) Search(word string) bool {
	var (
		curNode  = t
		nextNode *TrieNode
		ok       bool
	)

	for _, c := range word {
		nextNode, ok = curNode.children[c]
		if !ok {
			return false
		}

		curNode = nextNode
	}

	return curNode.isWord
}

func (t *TrieNode) Remove(word string) bool {
	if len(word) == 0 {
		return false
	}

	var (
		curNode        = t
		nextNode       *TrieNode
		ok             bool
		deleteToken    = rune(word[0])
		depthDeleteIdx int
	)

	for i, c := range word {
		if (curNode.isWord && !curNode.isLeaf()) || len(curNode.children) > 1 {
			depthDeleteIdx = i
			deleteToken = c
		}

		nextNode, ok = curNode.children[c]
		if !ok {
			return false
		}

		curNode = nextNode
	}

	if !curNode.isLeaf() || !curNode.isWord {
		return false
	}

	curNode = t
	for i := 0; i < depthDeleteIdx; i++ {
		curNode = curNode.children[rune(word[i])]
	}

	delete(curNode.children, deleteToken)
	return true
}

func (t *TrieNode) isLeaf() bool {
	return len(t.children) == 0
}
