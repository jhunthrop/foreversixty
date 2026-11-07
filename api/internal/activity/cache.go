package activity

import (
	"sync"
	"time"
)

// pageCache keeps the anonymous feed answers a short while, so a burst of signed-out home
// loads reads the table once per query per ttl. Entries are capped; a full cache is cleared.
type pageCache struct {
	mu      sync.Mutex
	ttl     time.Duration
	max     int
	entries map[string]cachedPage
}

type cachedPage struct {
	page Page
	at   time.Time
}

func newPageCache(ttl time.Duration, max int) *pageCache {
	return &pageCache{ttl: ttl, max: max, entries: map[string]cachedPage{}}
}

func (c *pageCache) get(key string, now time.Time) (Page, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.entries[key]
	if !ok || now.Sub(e.at) > c.ttl {
		return Page{}, false
	}
	return e.page, true
}

func (c *pageCache) put(key string, p Page, now time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.entries) >= c.max {
		c.entries = map[string]cachedPage{}
	}
	c.entries[key] = cachedPage{page: p, at: now}
}
