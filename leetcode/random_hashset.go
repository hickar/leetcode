package leetcode

import "math/rand"

/*
380. Insert Delete GetRandom O(1) (Medium)

Implement the RandomizedSet class:
RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present.
Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present.
Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements
(it's guaranteed that at least one element exists when this method is called).
Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average O(1) time complexity.
*/

const loadFactor float32 = 0.75

type RandomizedSet struct {
	items      [][]entry
	itemsCount int
	bufferSize int
}

func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		items:      make([][]entry, 1),
		bufferSize: 1,
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if s.shouldRehash() {
		s.rehash()
	}

	hash := calculateHash(val, s.bufferSize)

	switch {
	case s.items[hash] == nil:
		s.items[hash] = []entry{{Value: val, IsUsed: true}}
		return true
	default:
		for _, item := range s.items[hash] {
			if item.IsUsed && item.Value == val {
				return false
			}
		}
	}

	s.items[hash] = append(s.items[hash], entry{Value: val, IsUsed: true})
	s.itemsCount += 1
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	hash := calculateHash(val, s.bufferSize)

	for i, item := range s.items[hash] {
		if item.IsUsed && item.Value == val {
			s.items[hash] = append(s.items[hash][:i], s.items[hash][i+1:]...)
			s.itemsCount -= 1
			return true
		}
	}

	return false
}

func (s *RandomizedSet) shouldRehash() bool {
	return s.itemsCount+1 > int(float32(s.bufferSize)*loadFactor)
}

func (s *RandomizedSet) rehash() {
	newBufferSize := s.bufferSize * 2
	newItems := make([][]entry, newBufferSize)

	for _, bucket := range s.items {
		for _, item := range bucket {
			if item.IsUsed {
				newHash := calculateHash(item.Value, newBufferSize)
				newItems[newHash] = append(newItems[newHash], item)
			}
		}
	}

	s.items = newItems
	s.bufferSize = newBufferSize
}

func (s *RandomizedSet) GetRandom() int {
	for {
		randomHash := rand.Int() % s.bufferSize
		if len(s.items[randomHash]) == 0 {
			continue
		}

		randomItemIdx := rand.Int() % len(s.items[randomHash])
		return s.items[randomHash][randomItemIdx].Value
	}
}

func calculateHash(value int, limit int) int {
	return int(uint(value) % uint(limit))
}

type entry struct {
	Value  int
	IsUsed bool
}
