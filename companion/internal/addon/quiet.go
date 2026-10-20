package addon

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

const (
	// QuietFolder is how long nothing under an account's SavedVariables
	// folder may have changed before the companion writes into it.
	QuietFolder = 20 * time.Second
	// QuietSaved is how long ForeverSixty.lua, the file the game writes
	// at a reload or logout, must have been still. The client is busy
	// around that moment, so the inbox waits well past it.
	QuietSaved = 60 * time.Second
	// BusyRetry is how long a write waits after the game held the inbox
	// open when the rename was tried.
	BusyRetry = 30 * time.Second
	// BusyTries is how many held renames are waited out before the
	// companion copies over the file instead.
	BusyTries = 5
)

// pendingWrite is an inbox write that is due: since is when it first
// became due, retryAt the earliest next attempt after a held rename.
type pendingWrite struct {
	since   time.Time
	retryAt time.Time
	busy    int
}

// quiet reports whether the folder holding the SavedVariables file p
// has been still long enough to write into at now. Files the companion
// itself writes (the inbox and its temporary) do not count.
func quiet(p string, now time.Time) bool {
	dir := filepath.Dir(p)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return true // no folder, nothing for the game to be writing
	}
	inbox := filepath.Base(InboxPath(p))
	for _, e := range entries {
		if e.IsDir() || e.Name() == inbox || e.Name() == inbox+".tmp" {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		window := QuietFolder
		if e.Name() == filepath.Base(p) {
			window = QuietSaved
		}
		if now.Sub(info.ModTime()) < window {
			return false
		}
	}
	return true
}

// flush writes the inbox into every folder whose write is due and
// whose SavedVariables have been quiet; the rest wait for a later pass.
func (s *Sync) flush(now time.Time, paths []string) error {
	var errs []error
	for _, p := range paths {
		w := s.pending[p]
		if w == nil || now.Before(w.retryAt) || !s.rendered || !quiet(p, now) {
			continue
		}
		ip := InboxPath(p)
		wrote, err := WriteInbox(ip, s.body)
		if errors.Is(err, ErrBusy) {
			w.busy++
			if w.busy >= BusyTries {
				s.o.Log.Info("the game kept the addon inbox open; copying over it",
					"component", "addon", "path", ip, "tries", w.busy)
				wrote, err = WriteInboxCopy(ip, s.body)
			} else {
				w.retryAt = now.Add(BusyRetry)
				s.o.Log.Info("the game has the addon inbox open; trying again later",
					"component", "addon", "path", ip, "tries", w.busy)
				continue
			}
		}
		if err != nil {
			errs = append(errs, err)
			delete(s.pending, p)
			continue
		}
		deferred := now.Sub(w.since)
		delete(s.pending, p)
		if wrote {
			args := []any{"component", "addon", "path", ip,
				"builds", s.counts[0], "messages", s.counts[1]}
			if deferred > 0 {
				args = append(args, "deferred_ms", deferred.Milliseconds())
			}
			s.o.Log.Info("wrote the addon inbox", args...)
		}
	}
	return errors.Join(errs...)
}
