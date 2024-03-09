package leetcode

/*
146. LRU Cache (Medium)
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists.
	Otherwise, add the key-value pair to the cache.
If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.


Example 1:
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4


Constraints:
- 1 <= capacity <= 3000
- 0 <= key <= 104
- 0 <= value <= 105
- At most 2 * 105 calls will be made to get and put.
*/

type LRUCache struct {
	cap  int
	data map[int]*lruNode
	head *lruNode
	tail *lruNode
}

type lruNode struct {
	key   int
	value int
	prev  *lruNode
	next  *lruNode
}

func NewLRUCache(capacity int) LRUCache {
	head := &lruNode{}
	tail := &lruNode{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		data: make(map[int]*lruNode),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}

	this.push(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.data[key]
	if ok {
		node.value = value
		this.data[node.key] = node

		this.push(node)
		return
	}

	node = &lruNode{
		key:   key,
		value: value,
	}
	if len(this.data)+1 > this.cap {
		delete(this.data, this.tail.prev.key)
		this.pop()
	}

	this.data[node.key] = node
	this.addNode(node)
}

func (this *LRUCache) addNode(node *lruNode) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) push(node *lruNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) removeNode(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) pop() {
	this.removeNode(this.tail.prev)
}
