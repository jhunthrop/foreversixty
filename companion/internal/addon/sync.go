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
	"slices"
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
	// uploaded is the version of each SavedVariables file whose
	// exports the API has accepted. It is the sync's own record
	// rather than a question put to the shared cache, because the
	// status page reads through that cache too and would otherwise
	// consume the change the sync is waiting for. An upload that
	// failed is simply not recorded, so the next pass retries it.
	uploaded map[string]Stamp
	// builds and body are the inbox as it was last rendered. The
	// render is stamped with the time the builds changed, not the
	// time the file is written, so an unchanged inbox is the same
	// bytes pass after pass and WriteInbox's identity check fires.
	builds   []Build
	body     []byte
	rendered bool
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
	return &Sync{o: o, uploaded: map[string]Stamp{}}
}

// Poll does one pass. It is called on the same ticker as the pipeline
// and does nothing expensive when nothing has changed.
func (s *Sync) Poll(ctx context.Context, now time.Time) error {
	var errs []error
	paths := s.o.Paths()
	for _, p := range paths {
		exports, at, err := s.o.Cache.Exports(p)
		if errors.Is(err, fs.ErrNotExist) {
			continue // the player has not installed the addon here
		}
		if err != nil {
			s.o.Log.Warn("could not read the addon's saved variables",
				"component", "addon", "path", p, "err", err.Error())
			continue
		}
		if len(exports) == 0 {
			continue
		}
		if was, ok := s.uploaded[p]; ok && was.eq(at) {
			continue // this version is already on the site
		}
		if err := s.o.API.PostAddonExports(ctx, exports); err != nil {
			errs = append(errs, err) // unrecorded, so the next pass retries
			continue
		}
		s.uploaded[p] = at
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
	if !s.rendered || !slices.Equal(s.builds, inbox.Builds) {
		s.builds = slices.Clone(inbox.Builds)
		s.body, s.rendered = RenderInbox(inbox, now), true
	}
	for _, p := range paths {
		wrote, err := WriteInbox(InboxPath(p), s.body)
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
