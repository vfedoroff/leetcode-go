package solution

// Source: https://girai.dev/blog/lru-cache-implementation-in-go/

import (
	"container/list"
	"testing"
)

type LRUCache struct {
	cap int                   // capacity
	l   *list.List            // doubly linked list
	m   map[int]*list.Element // hash table for checking if list node exists
}

// Pair is the value of a list node.
type ListItem struct {
	key   int
	value int
}

// Constructor initializes a new LRUCache.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get a list node from the hash map.
func (c *LRUCache) Get(key int) int {
	// check if list node exists
	node, ok := c.m[key]
	if !ok {
		return -1
	}
	c.l.MoveToFront(node)
	return node.Value.(*list.Element).Value.(ListItem).value
}

// Put key and value in the LRUCache
func (c *LRUCache) Put(key int, value int) {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value.(*list.Element).Value = ListItem{key: key, value: value}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(ListItem).key
			// delete the node pointer in the hash map by key
			delete(c.m, idx)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := &list.Element{
			Value: ListItem{
				key:   key,
				value: value,
			},
		}
		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
		c.m[key] = ptr
	}
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // returns 1
	cache.Put(3, 3) // evicts key 2
	cache.Get(2)    // returns -1 (not found)
	cache.Put(4, 4) // evicts key 1
	cache.Get(1)    // returns -1 (not found)
	cache.Get(3)    // returns 3
	cache.Get(4)    // returns 4
}
