package sheet

import (
	"container/list"
	"sync"
)

// lru is a byte-slice cache bounded by entry count and total bytes, safe for concurrent use.
type lru struct {
	mu         sync.Mutex
	maxEntries int
	maxBytes   int
	bytes      int
	order      *list.List // front is the most recently used
	items      map[string]*list.Element
}

type lruEntry struct {
	key   string
	value []byte
}

func newLRU(maxEntries, maxBytes int) *lru {
	return &lru{maxEntries: maxEntries, maxBytes: maxBytes, order: list.New(), items: map[string]*list.Element{}}
}

func (c *lru) get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.order.MoveToFront(el)
	return el.Value.(*lruEntry).value, true
}

// put stores value under key, evicting the least recently used entries past either bound. A
// value larger than maxBytes on its own is not kept.
func (c *lru) put(key string, value []byte) {
	if c.maxBytes > 0 && len(value) > c.maxBytes {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.bytes += len(value) - len(el.Value.(*lruEntry).value)
		el.Value = &lruEntry{key: key, value: value}
		c.order.MoveToFront(el)
	} else {
		c.items[key] = c.order.PushFront(&lruEntry{key: key, value: value})
		c.bytes += len(value)
	}
	for c.order.Len() > c.maxEntries || (c.maxBytes > 0 && c.bytes > c.maxBytes) {
		last := c.order.Back()
		e := last.Value.(*lruEntry)
		c.order.Remove(last)
		delete(c.items, e.key)
		c.bytes -= len(e.value)
	}
}

func (c *lru) len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.order.Len()
}
