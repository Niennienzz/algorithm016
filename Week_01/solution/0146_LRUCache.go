package solution

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Could you do get and put in O(1) time complexity?
//
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists.
// Otherwise, add the key-value pair to the cache.
// If the number of keys exceeds the capacity from this operation, evict the least recently used key.
//
// Examples:
//
// Input
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// Explanation
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // cache is {1=1}
// lRUCache.put(2, 2); // cache is {1=1, 2=2}
// lRUCache.get(1);    // return 1
// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
// lRUCache.get(2);    // returns -1 (not found)
// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
// lRUCache.get(1);    // return -1 (not found)
// lRUCache.get(3);    // return 3
// lRUCache.get(4);    // return 4

type LRUCache struct {
	head     *cacheNode
	tail     *cacheNode
	values   map[int]*cacheNode
	capacity int
}

func NewLRUCache(capacity int) LRUCache {
	head := newCacheNode(-1, -1)
	tail := newCacheNode(-1, -1)

	head.next = tail
	tail.prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		values:   make(map[int]*cacheNode),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.values[key]; ok {
		node := this.values[key]
		this.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, val int) {
	// Update existing key.
	if node, ok := this.values[key]; ok {
		this.moveToHead(node)
		node.val = val
		return
	}

	// Check capacity, add new key.
	if len(this.values) == this.capacity {
		node := this.popTail()
		delete(this.values, node.key)
		node.key = key
		node.val = val
		this.addNode(node)
		this.values[key] = node
	} else {
		node := newCacheNode(key, val)
		this.addNode(node)
		this.values[key] = node
	}
}

// Always add right after head.
func (this *LRUCache) addNode(node *cacheNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *cacheNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) moveToHead(node *cacheNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *cacheNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// A doubly-linked-list node.
type cacheNode struct {
	key  int
	val  int
	prev *cacheNode
	next *cacheNode
}

func newCacheNode(key, val int) *cacheNode {
	return &cacheNode{
		key:  key,
		val:  val,
		prev: nil,
		next: nil,
	}
}
