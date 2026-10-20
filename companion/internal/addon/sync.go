// companion/internal/addon/sync.go
// The addon sync loop: watch each SavedVariables file's modification
// time, upload its exports and its outbox when the player logs out,
// and write the inbox back on a timer.
package addon

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"reflect"
	"slices"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

// InboxEvery is how often the inbox is refreshed, per the contract.
const InboxEvery = 10 * time.Minute

// API is the part of the client the sync uses.
type API interface {
	// PostAddonExports sends the character exports and, when outbox is
	// not nil, the addon's outbox in the same request.
	PostAddonExports(ctx context.Context, chars []Export, outbox *Outbox) (ExportsResult, error)
	AddonInbox(ctx context.Context) (Inbox, error)
}

// refusal is how the sync tells a site that read the outbox and said
// no from a request that never arrived: client.Error answers it. A
// refused outbox is never sent again; a lost one is, on the next pass.
type refusal interface {
	Refuses(field string) bool
}

// reasoned is an error that can say, in the site's own words, why it
// was refused: client.Error's message.
type reasoned interface {
	Reason() string
}

// reasonOf is the site's reason for err, or the error's text when the
// site gave none.
func reasonOf(err error) string {
	var r reasoned
	if errors.As(err, &r) {
		if why := r.Reason(); why != "" {
			return why
		}
	}
	return err.Error()
}

// refused reports whether err is the site refusing one part of the
// body: "outbox" (the outbox cannot be used) or "characters" (the
// character list was the wrong shape). Neither is worth sending again
// unchanged.
func refused(err error, field string) bool {
	var r refusal
	return errors.As(err, &r) && r.Refuses(field)
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
	// StateDir is the companion's state directory, where the keys of
	// the outboxes the site has answered are kept so a restart never
	// sends one twice (state.LoadOutboxSent). Empty keeps them in
	// memory only.
	StateDir string
	// Now is the clock the status page's "last sync" is stamped with;
	// nil is time.Now.
	Now func() time.Time
}

// Status is what the status page shows of the sync: when the site last
// saw an export (the newest ack's export_seen_at across characters),
// and what became of the newest in-game sim request, in the ack's
// words (SimQueued and the rest). Both are zero until the site acks.
type Status struct {
	SeenAt          time.Time
	SimRequestState string
	// Sent is the newest character list the site accepted: how many
	// characters it held and when it went up. Zero until one has.
	SentCount int
	SentAt    time.Time
	// Refusal is the site's reason when it refused the newest list
	// (a 400 naming the characters), and RefusedAt when. It is empty
	// again once a later list is accepted.
	Refusal   string
	RefusedAt time.Time
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
	// sent is, per SavedVariables file, the key of the last outbox the
	// site answered (stored or refused) or the companion dropped as too
	// large: a file the game rewrites with the same outbox sends its
	// exports again but not the outbox. It is kept in Options.StateDir,
	// so a restart does not send it again either.
	sent map[string]string
	// skipLogged and formatLogged hold what has been logged once this
	// process: a skipped inbox message by (type, position), an inbox
	// format this build does not render.
	skipLogged   map[string]bool
	formatLogged map[int]bool
	// builds, messages and body are the inbox as it was last rendered.
	// The render is stamped with the time the content changed, not the
	// time the file is written, so an unchanged inbox is the same
	// bytes pass after pass and WriteInbox's identity check fires.
	builds   []Build
	messages []Message
	body     []byte
	rendered bool

	// mu guards status, which the status page reads from the loopback
	// server's goroutine while Poll runs on the ticker's.
	mu     sync.Mutex
	status Status
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
	if o.Now == nil {
		o.Now = time.Now
	}
	s := &Sync{o: o, uploaded: map[string]Stamp{}, sent: map[string]string{},
		skipLogged: map[string]bool{}, formatLogged: map[int]bool{}}
	if o.StateDir != "" {
		sent, err := state.LoadOutboxSent(o.StateDir)
		if err != nil {
			// The worst a lost record costs is one outbox sent twice,
			// which the site answers as a duplicate.
			o.Log.Warn("could not read which outboxes the site has answered; starting over",
				"component", "addon", "err", err.Error())
		}
		if sent != nil {
			s.sent = sent
		}
		last, err := state.LoadLastSync(o.StateDir)
		if err != nil {
			o.Log.Warn("could not read when the characters last went up; starting over",
				"component", "addon", "err", err.Error())
		}
		s.status.SentCount, s.status.SentAt = last.Count, last.At
		s.status.Refusal, s.status.RefusedAt = last.Refusal, last.RefusedAt
	}
	return s
}

// markSent records that the site has answered (or the companion
// dropped) the outbox with key for file p, in memory and in the state
// directory.
func (s *Sync) markSent(p, key string) {
	s.sent[p] = key
	if s.o.StateDir == "" {
		return
	}
	if err := state.SaveOutboxSent(s.o.StateDir, s.sent); err != nil {
		s.o.Log.Warn("could not record which outboxes the site has answered",
			"component", "addon", "err", err.Error())
	}
}

// Status is the sync's latest word from the site.
func (s *Sync) Status() Status {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.status
}

// Poll does one pass. It is called on the same ticker as the pipeline
// and does nothing expensive when nothing has changed.
func (s *Sync) Poll(ctx context.Context, now time.Time) error {
	var errs []error
	paths := s.o.Paths()
	// outboxUp is set when an outbox reached the site this pass: its
	// ack is waiting, so the inbox is fetched now rather than when the
	// ten-minute timer comes round, and the player sees "Seen by the
	// site" within seconds of logging out.
	outboxUp := false
	for _, p := range paths {
		sent, err := s.upload(ctx, p)
		if err != nil {
			errs = append(errs, err)
		}
		outboxUp = outboxUp || sent
	}

	if !outboxUp && !s.lastInbox.IsZero() && now.Sub(s.lastInbox) < s.o.Every {
		return errors.Join(errs...)
	}
	inbox, err := s.o.API.AddonInbox(ctx)
	if err != nil {
		return errors.Join(append(errs, err)...)
	}
	s.lastInbox = now
	if inbox.Format != 0 && inbox.Format != InboxFormat {
		// A format this build does not know: render nothing from it and
		// leave the file the addon has. Format 0 is an API that predates
		// the field, whose body this build reads as it always has.
		if !s.formatLogged[inbox.Format] {
			s.formatLogged[inbox.Format] = true
			s.o.Log.Warn("the site sent an inbox format this companion does not know; "+
				"the addon keeps the inbox it has until the companion is updated",
				"component", "addon", "format", inbox.Format)
		}
		return errors.Join(errs...)
	}
	for _, sk := range inbox.Skipped {
		key := fmt.Sprintf("%s\x00%d", sk.Type, sk.Position)
		if s.skipLogged[key] {
			continue
		}
		s.skipLogged[key] = true
		s.o.Log.Warn("skipped an inbox message the addon could not read", "component", "addon",
			"type", sk.Type, "position", sk.Position, "why", sk.Why)
	}
	s.acked(inbox.Messages)
	// Message carries slices and maps of its own, which are not
	// comparable with slices.Equal's default ==, so the messages half
	// of this check is a deep-equal rather than the builds half's
	// cheap one; the inbox is small (InboxLimit builds plus a handful
	// of messages) and this runs once every ten minutes, not per poll.
	if !s.rendered || !slices.Equal(s.builds, inbox.Builds) || !reflect.DeepEqual(s.messages, inbox.Messages) {
		s.builds = slices.Clone(inbox.Builds)
		s.messages = slices.Clone(inbox.Messages)
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
				"path", InboxPath(p), "builds", len(inbox.Builds), "messages", len(inbox.Messages))
		}
	}
	return errors.Join(errs...)
}

// upload sends one SavedVariables file's exports, with its outbox when
// the outbox is one the site has not answered yet, once per version of
// the file. The outbox never goes alone: it rides the first post, with
// its own character in it, so the site can match the outbox's export
// to a character by name; any characters that do not fit go in posts
// of their own after it. It reports whether an outbox reached the site.
func (s *Sync) upload(ctx context.Context, p string) (bool, error) {
	saved, at, err := s.o.Cache.Saved(p)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil // the player has not installed the addon here
	}
	if err != nil {
		s.o.Log.Warn("could not read the addon's saved variables",
			"component", "addon", "path", p, "err", err.Error())
		return false, nil
	}
	if len(saved.Exports) == 0 {
		return false, nil
	}
	if was, ok := s.uploaded[p]; ok && was.eq(at) {
		return false, nil // this version is already on the site
	}
	ob := saved.Outbox
	if ob != nil && s.sent[p] == ob.key() {
		ob = nil // the site has answered this outbox already
	}
	// dropped is an outbox too large to send: on its own (MaxOutboxBytes)
	// or with even the one character it must ride with (MaxExportsBody).
	// It is answered, and said so, once the characters are up without it.
	var dropped *Outbox
	var first []Export
	var rest [][]Export
	if ob != nil {
		fits := len(ob.Body) <= MaxOutboxBytes
		if fits {
			if first, rest, fits, err = outboxBatches(saved.Exports, ob); err != nil {
				return false, err
			}
		}
		if !fits {
			dropped, ob = ob, nil
		}
	}

	outboxUp := false
	if ob != nil {
		var res ExportsResult
		res, err = s.o.API.PostAddonExports(ctx, first, ob)
		switch {
		case err == nil:
			// The outbox is on the site whatever becomes of the posts
			// after it, so it is answered now and never sent again.
			outboxUp = true
			s.markSent(p, ob.key())
			s.answered(res)
			err = s.postBatches(ctx, rest)
		case refused(err, "outbox"):
			// The site read the outbox and cannot use it. Sending it again
			// would be refused again, and must not hold the exports back.
			s.o.Log.Warn(CopyOutboxRefused, "component", "addon", "path", p, "err", err.Error())
			s.markSent(p, ob.key())
			ob = nil
			err = s.postCharacters(ctx, saved.Exports)
		}
	} else {
		err = s.postCharacters(ctx, saved.Exports)
	}
	if refused(err, "characters") {
		// The site read the list and cannot use it: the identical body
		// would be refused again on every pass. This version of the file
		// is done with; the next logout writes a new one.
		s.o.Log.Warn(CopyCharactersRefused, "component", "addon", "path", p, "err", err.Error())
		s.refusedList(reasonOf(err))
		s.uploaded[p] = at
		return outboxUp, nil
	}
	if err != nil {
		return outboxUp, err // unrecorded, so the next pass retries what is not yet answered
	}
	s.uploaded[p] = at
	s.o.Log.Info("uploaded character exports", "component", "addon",
		"path", p, "characters", len(saved.Exports), "outbox", ob != nil)
	s.sentList(len(saved.Exports))
	if dropped != nil {
		// The smallest post that could have carried it: the outbox with
		// its own character (which it must ride with) and nothing else.
		body, _ := ExportsBody(ownFirst(saved.Exports, dropped)[:1], dropped)
		s.o.Log.Warn(CopyOutboxTooLarge, "component", "addon", "path", p,
			"outbox_bytes", len(dropped.Body), "smallest_body_bytes", len(body),
			"outbox_limit", MaxOutboxBytes, "body_limit", MaxExportsBody)
		s.markSent(p, dropped.key())
	}
	return outboxUp, nil
}

// sentList records that the site accepted a character list of n, in
// memory and in the state directory. It clears an earlier refusal: the
// newest word on the list is the good one.
func (s *Sync) sentList(n int) {
	s.mu.Lock()
	s.status.SentCount, s.status.SentAt = n, s.o.Now()
	s.status.Refusal, s.status.RefusedAt = "", time.Time{}
	last := s.lastSync()
	s.mu.Unlock()
	s.saveLastSync(last)
}

// refusedList records the site's reason for refusing a character
// list. What it accepted before stays: the page says both.
func (s *Sync) refusedList(reason string) {
	s.mu.Lock()
	s.status.Refusal, s.status.RefusedAt = reason, s.o.Now()
	last := s.lastSync()
	s.mu.Unlock()
	s.saveLastSync(last)
}

// lastSync is the status as the state file keeps it; mu is held.
func (s *Sync) lastSync() state.LastSync {
	return state.LastSync{Count: s.status.SentCount, At: s.status.SentAt,
		Refusal: s.status.Refusal, RefusedAt: s.status.RefusedAt}
}

func (s *Sync) saveLastSync(last state.LastSync) {
	if s.o.StateDir == "" {
		return
	}
	if err := state.SaveLastSync(s.o.StateDir, last); err != nil {
		s.o.Log.Warn("could not record when the characters last went up",
			"component", "addon", "err", err.Error())
	}
}

// answered logs the site's answer for an outbox, one plain line per
// part, and keeps the sim request's outcome for the status page.
func (s *Sync) answered(res ExportsResult) {
	if res.Outbox == nil {
		return // a site that predates the outbox stored the exports only
	}
	for _, line := range res.Outbox.Lines() {
		s.o.Log.Info(line, "component", "addon")
	}
	if st := SimStateOf(res.Outbox.SimRequest); st != "" {
		s.mu.Lock()
		s.status.SimRequestState = st
		s.mu.Unlock()
	}
}

// acked takes the status page's line from the inbox's acks: the newest
// export_seen_at across characters, and the sim request state of the
// newest ack that carries one. The ack is the site's own word, so it
// replaces what the export answer said; an inbox with no ack changes
// nothing.
func (s *Sync) acked(messages []Message) {
	var seen, simAt time.Time
	var state string
	for _, m := range messages {
		if m.Type != MessageAck {
			continue
		}
		at, err := time.Parse(time.RFC3339, m.ExportSeenAt)
		if err != nil {
			s.o.Log.Warn("an ack carried a time that cannot be read", "component", "addon",
				"export_seen_at", m.ExportSeenAt)
			continue
		}
		if at.After(seen) {
			seen = at
		}
		if m.SimRequestState != "" && (state == "" || at.After(simAt)) {
			state, simAt = m.SimRequestState, at
		}
	}
	if seen.IsZero() {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status.SeenAt = seen
	if state != "" {
		s.status.SimRequestState = state
	}
}

// postCharacters sends the character exports alone, in as many posts
// as MaxExportsBody and MaxExportsPerPost need. Any failed post fails
// the upload, so the next pass sends them all again; storing an export
// twice is an overwrite.
func (s *Sync) postCharacters(ctx context.Context, chars []Export) error {
	batches, err := exportBatches(chars)
	if err != nil {
		return err
	}
	return s.postBatches(ctx, batches)
}

// postBatches sends each batch as a characters-only post, in order,
// stopping at the first failure.
func (s *Sync) postBatches(ctx context.Context, batches [][]Export) error {
	for _, batch := range batches {
		if _, err := s.o.API.PostAddonExports(ctx, batch, nil); err != nil {
			return err
		}
	}
	return nil
}
