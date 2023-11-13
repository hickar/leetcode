package structures

import "encoding/binary"

const loadFactor float32 = 0.75

type HashSet[T Hashable] struct {
	itemCount  int
	bufferSize int

	items [][]entry[T]
}

type entry[T Hashable] struct {
	Value  T
	IsUsed bool
}

type Hashable interface {
	comparable
	Hash() []byte
}

func NewHashSet[T Hashable]() *HashSet[T] {
	return &HashSet[T]{
		items:      make([][]entry[T], 1),
		bufferSize: 1,
	}
}

func NewHashSetSized[T Hashable](size int) *HashSet[T] {
	return &HashSet[T]{
		items:      make([][]entry[T], size),
		bufferSize: 1,
	}
}

func (h *HashSet[T]) InsertMany(values ...T) {
	for _, value := range values {
		h.Insert(value)
	}
}

func (h *HashSet[T]) Insert(value T) bool {
	if h.shouldRehash() {
		h.rehash()
	}

	hash := calculateHash(value, h.bufferSize)
	switch {
	case h.items[hash] == nil:
		h.items[hash] = append(h.items[hash], entry[T]{Value: value, IsUsed: true})
	default:
		for _, item := range h.items[hash] {
			if item.IsUsed && item.Value == value {
				return false
			}
		}
	}

	h.items[hash] = append(h.items[hash], entry[T]{Value: value, IsUsed: true})
	h.itemCount += 1
	return true
}

func (h *HashSet[T]) Contains(key T) bool {
	hash := calculateHash(key, h.bufferSize)

	for _, item := range h.items[hash] {
		if item.IsUsed && item.Value == key {
			return true
		}
	}

	return false
}

func (h *HashSet[T]) Remove(key T) bool {
	hash := calculateHash(key, h.bufferSize)

	for i, item := range h.items[hash] {
		if item.IsUsed && item.Value == key {
			h.items[hash] = append(h.items[hash][:i], h.items[hash][i+1:]...)
			h.itemCount -= 1
			return true
		}
	}

	return false
}

func (h *HashSet[T]) rehash() {
	newBufferSize := max(h.bufferSize*2, 1)
	newItems := make([][]entry[T], newBufferSize)

	for _, bucket := range h.items {
		for _, item := range bucket {
			if !item.IsUsed {
				continue
			}

			newHash := calculateHash(item.Value, newBufferSize)
			newItems[newHash] = append(newItems[newHash], item)
		}
	}

	h.items = newItems
	h.bufferSize = newBufferSize
}

func (h *HashSet[T]) shouldRehash() bool {
	return h.itemCount+1 > int(float32(h.bufferSize)*loadFactor)
}

func (h *HashSet[T]) Len() int {
	return h.itemCount
}

func (h *HashSet[T]) Cap() int {
	return h.bufferSize
}

func (h *HashSet[T]) Slice() []T {
	s := make([]T, 0, h.itemCount)

	for _, bucket := range h.items {
		if bucket == nil {
			continue
		}

		for _, item := range bucket {
			if !item.IsUsed {
				continue
			}

			s = append(s, item.Value)
		}
	}

	return s
}

const (
	fnvOffsetBasis uint64 = 14695981039346656037
	fnvPrime       uint64 = 1099511628211
)

func calculateHash[T Hashable](v T, limit int) int {
	// Fowler–Noll–Vo hash function
	// https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function
	hash := fnvOffsetBasis
	for _, b := range v.Hash() {
		hash = hash ^ uint64(b)
		hash = hash * fnvPrime
	}

	return int(hash % uint64(limit))
}

type IntHashable int

func (i IntHashable) Hash() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(buf, int64(i))
	return buf
}

type StringHashable string

func (s StringHashable) Hash() []byte {
	return []byte(s)
}
