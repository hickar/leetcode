package leetcode

type Trie struct {
	nodes      map[byte]*Trie
	isComplete bool
}

func Constructor() Trie {
	return Trie{nodes: make(map[byte]*Trie)}
}

func (t *Trie) Insert(word string) {
	var ok bool
	cur := t

	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok = cur.nodes[c]; !ok {
			tr := Constructor()
			cur.nodes[c] = &tr
		}

		cur = cur.nodes[c]
	}

	cur.isComplete = true
}

func (t *Trie) Search(word string) bool {
	var ok bool
	cur := t

	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok = cur.nodes[c]; !ok {
			return false
		}

		cur = cur.nodes[c]
	}

	return cur.isComplete
}

func (t *Trie) StartsWith(prefix string) bool {
	var ok bool
	cur := t

	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok = cur.nodes[c]; !ok {
			return false
		}

		cur = cur.nodes[c]
	}

	return true
}
