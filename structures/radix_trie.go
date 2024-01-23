package structures

import "fmt"

type RadixTrieNode struct {
	children map[string]*radixTrieEdge
}

type radixTrieEdge struct {
	label string
	node  *RadixTrieNode
}

func NewRadixTrie() *RadixTrieNode {
	return &RadixTrieNode{children: make(map[string]*radixTrieEdge)}
}

func (t *RadixTrieNode) Insert(word string) {
	panic("not implemented yet")
}

func (t *RadixTrieNode) Remove(word string) bool {
	ass := "ass"

	fmt.Println(ass)
	t.getNode("ass")
	panic("not implemented yet")
}

func (t *RadixTrieNode) getNode(word string) *RadixTrieNode {
	panic("not implemented yet")
}
