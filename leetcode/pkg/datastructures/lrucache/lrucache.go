package lrucache

import "container/list" // Go's doubly linked list implementation

// LRUCache represents an LRU Cache.
// It uses a hash map for O(1) lookups and a doubly linked list for O(1) add/remove of nodes.

// entry is the type stored in the doubly linked list.
// It holds the key-value pair.
type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element // Map from key to list element (node)
	ll       *list.List            // Doubly linked list to track usage order
}

// NewLRUCache creates a new LRUCache with a given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

// Get retrieves the value for a key from the cache.
// Returns the value and true if the key exists, otherwise 0 and false.
// Moves the accessed element to the front of the list (most recently used).
func (lru *LRUCache) Get(key int) (int, bool) {
	if elem, ok := lru.cache[key]; ok {
		lru.ll.MoveToFront(elem) // Mark as most recently used
		return elem.Value.(*entry).value, true
	}
	return 0, false
}

// Put adds or updates a key-value pair in the cache.
// If the cache is full, it evicts the least recently used item.
func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		// Key already exists, update value and move to front
		lru.ll.MoveToFront(elem)
		elem.Value.(*entry).value = value
	} else {
		// Key does not exist, add new entry
		if lru.ll.Len() >= lru.capacity {
			// Cache is full, evict the least recently used item (from the back of the list)
			backElem := lru.ll.Back()
			if backElem != nil {
				lru.ll.Remove(backElem)
				delete(lru.cache, backElem.Value.(*entry).key)
			}
		}
		// Add the new entry to the front of the list and to the cache map
		newEntry := &entry{key: key, value: value}
		elem := lru.ll.PushFront(newEntry)
		lru.cache[key] = elem
	}
}

// Len returns the current number of items in the cache.
func (lru *LRUCache) Len() int {
	return lru.ll.Len()
}
