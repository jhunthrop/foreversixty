// companion/internal/addon/cache.go
// One reader for the SavedVariables files, memoised on modification
// time. A SavedVariables file is megabytes of Lua and changes only
// when the player logs out, while two callers want its contents: the
// sync, which uploads on a change, and the status page, which is
// polled every two seconds. Both go through the cache, so the file is
// parsed once per logout rather than thirty times a minute.
package addon

import (
	"os"
	"sync"
	"time"
)

// Cache remembers each SavedVariables file's exports until the file
// changes. The zero value is ready to use and is safe for concurrent
// callers, which it has: the sync runs on the companion's ticker and
// the status page on the loopback server's goroutine.
type Cache struct {
	mu   sync.Mutex
	seen map[string]cached
}

// cached is one file's last read, good until the file moves.
type cached struct {
	mod     time.Time
	size    int64
	exports []Export
	// err is remembered with the rest: a file the reader cannot parse
	// does not become parseable until the game writes it again, and
	// retrying the parse every two seconds is the cost this cache
	// exists to remove.
	err error
}

// Exports returns one file's character exports and whether this call
// re-read the file. A file that has not moved since the last call is
// answered from memory. A file that cannot be stat'ed is not cached:
// the error is the caller's to classify, and a missing file is the
// normal state where the addon is not installed.
func (c *Cache) Exports(path string) ([]Export, bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, false, err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if was, ok := c.seen[path]; ok && was.mod.Equal(fi.ModTime()) && was.size == fi.Size() {
		return was.exports, false, was.err
	}
	found, err := ScanExports(path)
	if c.seen == nil {
		c.seen = map[string]cached{}
	}
	c.seen[path] = cached{mod: fi.ModTime(), size: fi.Size(), exports: found, err: err}
	return found, true, err
}

// Forget drops what is remembered about one file, so the next read
// parses it again even though the file has not changed. The sync uses
// it to retry an upload the network refused.
func (c *Cache) Forget(path string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.seen, path)
}
