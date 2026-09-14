// companion/internal/addon/cache.go
// One reader for the SavedVariables files, memoised on modification
// time. A SavedVariables file is megabytes of Lua and changes only
// when the player logs out, while two callers want its contents: the
// sync, which uploads on a change, and the status page, which is
// polled every two seconds. Both go through the cache, so the file is
// parsed once per logout rather than thirty times a minute.
//
// The cache is a parse memo and nothing more. It does not tell a
// caller whether this call was the one that re-read the file: two
// callers share it and only one of them could ever win that answer,
// which is how the exports present at launch stopped being uploaded.
// A caller that must act once per change compares Stamps instead.
package addon

import (
	"os"
	"sync"
	"time"
)

// Stamp identifies one version of a file: the modification time and
// size the cache keys on. It is returned alongside the exports so a
// caller with work to do per change — the sync uploads once per
// logout — can remember the version it last acted on and compare,
// rather than ask the shared cache a question only one caller can be
// told the truth about.
type Stamp struct {
	Mod  time.Time
	Size int64
}

// eq reports whether two stamps name the same version of a file.
// time.Time is not compared with == because a wall clock reading
// carries a location the comparison has no opinion about.
func (s Stamp) eq(o Stamp) bool { return s.Size == o.Size && s.Mod.Equal(o.Mod) }

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
	at      Stamp
	exports []Export
	// err is remembered with the rest: a file the reader cannot parse
	// does not become parseable until the game writes it again, and
	// retrying the parse every two seconds is the cost this cache
	// exists to remove.
	err error
}

// Exports returns one file's character exports and the version of the
// file they were parsed from. A file that has not moved since the
// last call is answered from memory. A file that cannot be stat'ed is
// not cached: the error is the caller's to classify, and a missing
// file is the normal state where the addon is not installed.
func (c *Cache) Exports(path string) ([]Export, Stamp, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, Stamp{}, err
	}
	at := Stamp{Mod: fi.ModTime(), Size: fi.Size()}
	c.mu.Lock()
	defer c.mu.Unlock()
	if was, ok := c.seen[path]; ok && was.at.eq(at) {
		return was.exports, at, was.err
	}
	found, err := ScanExports(path)
	if c.seen == nil {
		c.seen = map[string]cached{}
	}
	c.seen[path] = cached{at: at, exports: found, err: err}
	return found, at, err
}
