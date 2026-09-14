// companion/internal/addon/sync.go
// The addon sync loop: watch each SavedVariables file's modification
// time, upload its exports when the player logs out, and write the
// inbox back on a timer.
package addon

import (
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"time"
)

// InboxEvery is how often the inbox is refreshed, per the contract.
const InboxEvery = 10 * time.Minute

// API is the part of the client the sync uses.
type API interface {
	PostAddonExports(ctx context.Context, chars []Export) error
	AddonInbox(ctx context.Context) (Inbox, error)
}

// Options configures a Sync.
type Options struct {
	// Paths returns the SavedVariables files to watch, one per
	// account folder per install. It is a function because the player
	// can add an install while the companion runs.
	Paths func() []string
	API   API
	Log   *slog.Logger
	Every time.Duration
	// Cache is the shared reader for the SavedVariables files. A nil
	// Cache gets one of its own; the companion passes the same one
	// the status page reads through.
	Cache *Cache
}

// Sync moves strings between the addon and the API.
type Sync struct {
	o         Options
	lastInbox time.Time
}

// New builds a sync.
func New(o Options) *Sync {
	if o.Every == 0 {
		o.Every = InboxEvery
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	if o.Cache == nil {
		o.Cache = &Cache{}
	}
	return &Sync{o: o}
}

// Poll does one pass. It is called on the same ticker as the pipeline
// and does nothing expensive when nothing has changed.
func (s *Sync) Poll(ctx context.Context, now time.Time) error {
	var errs []error
	paths := s.o.Paths()
	for _, p := range paths {
		exports, fresh, err := s.o.Cache.Exports(p)
		if errors.Is(err, fs.ErrNotExist) {
			continue // the player has not installed the addon here
		}
		if err != nil {
			s.o.Log.Warn("could not read the addon's saved variables",
				"component", "addon", "path", p, "err", err.Error())
			continue
		}
		if !fresh || len(exports) == 0 {
			continue
		}
		if err := s.o.API.PostAddonExports(ctx, exports); err != nil {
			s.o.Cache.Forget(p) // read it again, and retry, next pass
			errs = append(errs, err)
			continue
		}
		s.o.Log.Info("uploaded character exports", "component", "addon",
			"path", p, "characters", len(exports))
	}

	if !s.lastInbox.IsZero() && now.Sub(s.lastInbox) < s.o.Every {
		return errors.Join(errs...)
	}
	inbox, err := s.o.API.AddonInbox(ctx)
	if err != nil {
		return errors.Join(append(errs, err)...)
	}
	s.lastInbox = now
	body := RenderInbox(inbox, now)
	for _, p := range paths {
		wrote, err := WriteInbox(InboxPath(p), body)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if wrote {
			s.o.Log.Info("wrote the addon inbox", "component", "addon",
				"path", InboxPath(p), "builds", len(inbox.Builds))
		}
	}
	return errors.Join(errs...)
}
