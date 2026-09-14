// companion/internal/pipeline/pipeline.go
// Package pipeline is the live path: bytes from the watcher into the
// engine, closed fights and raw chunks into the queue, the queue into
// the API. Parsing and uploading are two separate calls — Tick and
// Drain — so a network outage stops the uploads and never the parse,
// and so a test can drive both by hand.
//
// Every offset the companion sends is relative to the report's own
// byte stream, which begins at zero when the report begins. The
// watcher works in file offsets; state.Report.StartOffset is the one
// place the two meet.
//
// One mutex guards the open report. Tick, Live and Drain all run on
// the companion's single ticker, but Status is read from the loopback
// UI's HTTP goroutine while that ticker is inside Tick, so the fields
// Status reads are locked. The lock never spans a network call.
package pipeline

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Defaults from the interface contract: raw is streamed in 4 MiB
// chunks and the live snapshot goes up every five seconds.
const (
	RawChunk  = 4 << 20
	LiveEvery = 5 * time.Second
	// StateEvery bounds how often the engine's session is serialised
	// to disk between fights. Fight closes always save.
	StateEvery = 30 * time.Second
)

// Options configures a Pipeline.
type Options struct {
	StateDir string
	Client   *client.Client
	Queue    *queue.Queue
	Watch    *watch.Watcher
	Config   config.Config
	Log      *slog.Logger

	RawChunk   int
	LiveEvery  time.Duration
	StateEvery time.Duration
	// NewKey mints a local report key. Tests pin it.
	NewKey func(time.Time) string
}

// Pipeline owns the open report.
type Pipeline struct {
	o   Options
	enc *zstd.Encoder

	// mu guards everything below it: the open report and the engine
	// session behind it, which Status reads from another goroutine.
	mu   sync.Mutex
	sess *session.Session
	cur  *state.Report
	// raw holds report-relative bytes not yet handed to the queue;
	// rawAt is the offset of raw[0].
	raw      []byte
	rawAt    int64
	lastLive time.Time
	lastSave time.Time
}

// New builds a pipeline and resumes the report that was open when the
// companion last stopped, if there is one.
func New(o Options) (*Pipeline, error) {
	if o.Client == nil || o.Queue == nil || o.Watch == nil {
		return nil, errors.New("pipeline: the client, queue and watcher are all required")
	}
	if o.RawChunk == 0 {
		o.RawChunk = RawChunk
	}
	if o.LiveEvery == 0 {
		o.LiveEvery = LiveEvery
	}
	if o.StateEvery == 0 {
		o.StateEvery = StateEvery
	}
	if o.NewKey == nil {
		o.NewKey = state.NewKey
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	enc, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	if err != nil {
		return nil, err
	}
	p := &Pipeline{o: o, enc: enc}
	if err := p.resume(); err != nil {
		return nil, err
	}
	return p, nil
}

// Close releases the compressor.
func (p *Pipeline) Close() error {
	p.enc.Close()
	return nil
}

// SetConfig swaps the configuration the pipeline reads when a report
// opens, so a visibility or a logging character chosen on the
// settings page is in force on the next report rather than on the
// next launch. A report already open keeps what it started with:
// state.Report captures both at the start on purpose.
func (p *Pipeline) SetConfig(c config.Config) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.o.Config = c
}

// SetWatch repoints the tail at another Logs directory and reports
// whether it did. It refuses while a report is open, because the
// replacement is a fresh watcher: it has seen no file, so it emits no
// Complete for the report already open, and the watcher that would
// have emitted one is the one being thrown away. p.cur would be left
// open with nobody left to close it, until the first Start from the
// new directory overwrote it — losing the rest of the night and its
// completion both. The caller asks again on the next tick, and the
// swap happens when the report closes.
func (p *Pipeline) SetWatch(w *watch.Watcher) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur != nil {
		return false
	}
	p.o.Watch = w
	return true
}

// sessionOptions are the engine settings the companion parses with.
// KeepEvents is on because every closed fight is written as Parquet.
func sessionOptions(reportKey string, base time.Time) session.Options {
	o := session.Options{
		ReportID:   reportKey,
		Base:       base,
		Infer:      true,
		KeepEvents: true,
		Units:      units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:      fight.DefaultOptions(),
		Summary:    summary.DefaultOptions(),
	}
	o.Summary.SpecNames = units.RetailSpecName
	return o
}

// resume restores the report that was still open, if any.
func (p *Pipeline) resume() error {
	reports, err := state.List(p.o.StateDir)
	if err != nil {
		return err
	}
	for i := len(reports) - 1; i >= 0; i-- {
		r := reports[i]
		if r.Closed {
			continue
		}
		if len(r.Session) == 0 {
			// Nothing had closed when the companion stopped: read the
			// report again from its first byte. Re-sent bytes are
			// ignored by offset on both sides.
			if err := p.o.Watch.Resume(r.LogPath, r.StartOffset, r.LastAppend); err != nil {
				p.o.Log.Warn("could not reopen the log; starting a new report",
					"component", "pipeline", "report", r.Key, "err", err.Error())
				p.abandon(r)
				continue
			}
			p.sess, p.cur = session.New(sessionOptions(r.Key, r.StartedAt)), &r
			p.refillRaw(r)
			p.o.Log.Info("restarted a report from its first byte",
				"component", "pipeline", "report", r.Key)
			return nil
		}
		s, err := session.Restore(sessionOptions(r.Key, r.StartedAt), r.Session)
		if err != nil {
			p.o.Log.Warn("could not restore a report; starting a new one",
				"component", "pipeline", "report", r.Key, "err", err.Error())
			p.abandon(r)
			continue
		}
		replay, err := session.ReplayOffset(r.Session)
		if err != nil {
			return err
		}
		if err := p.o.Watch.Resume(r.LogPath, r.StartOffset+replay, r.LastAppend); err != nil {
			p.o.Log.Warn("could not resume the log file; starting a new report",
				"component", "pipeline", "report", r.Key, "err", err.Error())
			p.abandon(r)
			continue
		}
		p.sess, p.cur = s, &r
		p.refillRaw(r)
		p.o.Log.Info("resumed a report", "component", "pipeline", "report", r.Key,
			"report_id", r.ReportID, "replay_from", replay)
		return nil
	}
	return nil
}

// abandon marks a report the companion can no longer continue as
// closed, so the next startup does not try to resume it again. Its
// queued work is unaffected: the state file stays, and the uploader
// still finds the report id on it.
func (p *Pipeline) abandon(r state.Report) {
	r.Closed = true
	if err := state.Save(p.o.StateDir, r); err != nil {
		p.o.Log.Error("could not close an abandoned report's state", "component", "pipeline",
			"report", r.Key, "err", err.Error())
	}
}

// Tick reads whatever the game has written and turns it into queued
// work. It never talks to the network.
func (p *Pipeline) Tick(now time.Time) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	events, err := p.o.Watch.Poll(now)
	if err != nil {
		return err
	}
	for _, e := range events {
		switch e.Kind {
		case watch.Start:
			if err := p.start(e, now); err != nil {
				return err
			}
		case watch.Append:
			if err := p.append(e, now); err != nil {
				return err
			}
		case watch.Complete:
			if err := p.complete(now); err != nil {
				return err
			}
		}
	}
	if p.cur != nil && now.Sub(p.lastSave) >= p.o.StateEvery {
		if err := p.save(); err != nil {
			return err
		}
		p.lastSave = now
	}
	return nil
}

func (p *Pipeline) start(e watch.Event, now time.Time) error {
	key := p.o.NewKey(now)
	p.cur = &state.Report{
		Key: key, LogPath: e.Path, StartOffset: e.Offset,
		Visibility: p.o.Config.ReportVisibility, LoggingCharacter: p.o.Config.LoggingCharacter,
		StartedAt: now, LastAppend: now,
		EngineVersion: session.Version,
	}
	p.sess = session.New(sessionOptions(key, now))
	p.raw, p.rawAt, p.lastLive, p.lastSave = nil, 0, time.Time{}, now
	p.o.Log.Info("a report started", "component", "pipeline", "report", key,
		"path", e.Path, "file_offset", e.Offset)
	return p.save()
}

func (p *Pipeline) append(e watch.Event, now time.Time) error {
	if p.cur == nil {
		return nil // bytes with no open report; the watcher opens one first
	}
	rel := e.Offset - p.cur.StartOffset
	res, err := p.sess.Feed(e.Data, rel)
	if err != nil {
		return fmt.Errorf("feed %s at %d: %w", e.Path, rel, err)
	}
	p.bufferRaw(rel, e.Data)
	p.cur.Offset = p.sess.Offset()
	p.cur.LastAppend = now
	for _, c := range res.Closed {
		if err := p.enqueueFight(c); err != nil {
			return err
		}
	}
	if err := p.flushRaw(false); err != nil {
		return err
	}
	if len(res.Closed) > 0 {
		if err := p.save(); err != nil {
			return err
		}
		p.lastSave = now
	}
	return nil
}

// bufferRaw appends the part of a chunk that has not been buffered
// already. A restart re-feeds bytes the engine needs to rebuild the
// open fight, and those bytes must not be uploaded twice.
func (p *Pipeline) bufferRaw(rel int64, data []byte) {
	next := p.rawAt + int64(len(p.raw))
	end := rel + int64(len(data))
	if end <= next {
		return
	}
	if rel < next {
		data = data[next-rel:]
		rel = next
	}
	if len(p.raw) == 0 {
		p.rawAt = rel
	}
	p.raw = append(p.raw, data...)
}

// refillRaw rebuilds the unqueued tail of the raw stream after a
// restart. Without it the bytes between the last queued chunk and the
// stopping point would never be uploaded, and the server could not
// reassemble the log.
func (p *Pipeline) refillRaw(r state.Report) {
	p.rawAt, p.raw = r.RawSent, nil
	if r.Offset <= r.RawSent {
		return
	}
	buf, err := watch.ReadRange(r.LogPath, r.StartOffset+r.RawSent, r.StartOffset+r.Offset)
	if err != nil {
		p.o.Log.Warn("could not reread the unsent raw bytes", "component", "pipeline",
			"report", r.Key, "err", err.Error())
		return
	}
	p.raw = buf
}

// flushRaw queues whole chunks, and on completion the remainder too.
func (p *Pipeline) flushRaw(final bool) error {
	for len(p.raw) >= p.o.RawChunk || (final && len(p.raw) > 0) {
		n := min(len(p.raw), p.o.RawChunk)
		// The digest is over the plaintext, because that is what the
		// server hashes after it decodes; it is taken here, while the
		// plaintext still exists.
		digest := client.SHA256(p.raw[:n])
		packed := p.enc.EncodeAll(p.raw[:n], nil)
		if _, err := p.o.Queue.Enqueue(queue.Item{
			Kind: queue.Raw, ReportKey: p.cur.Key, Offset: p.rawAt, SHA256: digest,
		}, packed); err != nil {
			return err
		}
		p.rawAt += int64(n)
		p.cur.RawSent = p.rawAt
		p.raw = p.raw[n:]
	}
	if len(p.raw) == 0 {
		p.raw = nil
	}
	return nil
}

// enqueueFight builds one closed fight's bundle and queues it.
func (p *Pipeline) enqueueFight(c session.Closed) error {
	events, err := parquet.Marshal(c.Events)
	if err != nil {
		return fmt.Errorf("write fight %d as parquet: %w", c.Fight.Index, err)
	}
	b := client.FightBundle{
		Summary: c.Summary,
		Events:  events,
		Metrics: client.MetricsRowsOf(c.Fight, c.Summary),
		RawRange: client.RawRange{
			StartOffset: c.Fight.StartOffset,
			EndOffset:   c.Fight.EndOffset,
			SHA256:      p.rawHash(c.Fight),
		},
	}
	ct, body, err := b.Encode()
	if err != nil {
		return err
	}
	if _, err := p.o.Queue.Enqueue(queue.Item{
		Kind: queue.Fight, ReportKey: p.cur.Key,
		FightIndex: c.Fight.Index, ContentType: ct,
	}, body); err != nil {
		return err
	}
	p.cur.Fights++
	if p.cur.Zone == "" && c.Fight.Zone != "" {
		p.cur.Zone = c.Fight.Zone
	}
	p.o.Log.Info("a fight closed", "component", "pipeline", "report", p.cur.Key,
		"fight", c.Fight.Index, "name", c.Fight.Name, "bundle_bytes", len(body))
	return nil
}

// rawHash is the hash of the fight's bytes, read back out of the log.
// A log that has already rotated away yields no hash rather than no
// fight: the server's later raw-sample check is what the hash serves,
// and a missing hash costs that check, not the report.
func (p *Pipeline) rawHash(f fight.Fight) string {
	raw, err := watch.ReadRange(p.cur.LogPath,
		p.cur.StartOffset+f.StartOffset, p.cur.StartOffset+f.EndOffset)
	if err != nil {
		p.o.Log.Warn("could not hash a fight's raw bytes", "component", "pipeline",
			"report", p.cur.Key, "fight", f.Index, "err", err.Error())
		return ""
	}
	return client.SHA256(raw)
}

// complete closes the engine's session, queues the last fights, the
// last raw chunk and the completion, and forgets the open report.
func (p *Pipeline) complete(now time.Time) error {
	if p.cur == nil {
		return nil
	}
	res, err := p.sess.Close()
	if err != nil {
		return fmt.Errorf("close the session for %s: %w", p.cur.Key, err)
	}
	for _, c := range res.Closed {
		if err := p.enqueueFight(c); err != nil {
			return err
		}
	}
	p.cur.Offset = p.sess.Offset()
	if err := p.flushRaw(true); err != nil {
		return err
	}
	body, err := marshalComplete(client.Complete{
		FinalOffset:   p.cur.Offset,
		EngineVersion: session.Version,
		Health:        p.sess.Health(),
	})
	if err != nil {
		return err
	}
	if _, err := p.o.Queue.Enqueue(queue.Item{
		Kind: queue.Complete, ReportKey: p.cur.Key,
	}, body); err != nil {
		return err
	}
	p.cur.Closed = true
	p.cur.Session = nil
	if err := p.save(); err != nil {
		return err
	}
	p.o.Log.Info("a report completed", "component", "pipeline", "report", p.cur.Key,
		"fights", p.cur.Fights, "final_offset", p.cur.Offset)
	p.sess, p.cur, p.raw = nil, nil, nil
	return nil
}

// save writes the open report's state, including the engine's session
// so a restart resumes rather than restarts.
func (p *Pipeline) save() error {
	if p.cur == nil {
		return nil
	}
	if p.sess != nil && !p.cur.Closed {
		// A session that has not settled a layout yet cannot be
		// serialised, and does not need to be: nothing has closed,
		// so a restart re-feeds the report from its first byte.
		if blob, err := p.sess.State(); err == nil {
			p.cur.Session = blob
		} else {
			p.cur.Session = nil
			p.o.Log.Debug("the session is not resumable yet", "component", "pipeline",
				"report", p.cur.Key, "err", err.Error())
		}
	}
	return state.Save(p.o.StateDir, *p.cur)
}

// Live uploads the running summary of the fight in progress. It is not
// queued: the next snapshot replaces a failed one. It is silent when
// no fight is open or the report has no server id yet.
func (p *Pipeline) Live(ctx context.Context, now time.Time) error {
	reportID, f, sum, due := p.liveDue(now)
	if !due {
		return nil
	}
	err := p.o.Client.PutLive(ctx, reportID, f.Index, client.Live{
		Summary:   sum,
		ElapsedMS: now.Sub(f.Start).Milliseconds(),
		UpdatedAt: now.UTC(),
	})
	if err != nil {
		p.o.Log.Debug("a live snapshot did not land", "component", "pipeline",
			"report", reportID, "fight", f.Index, "err", err.Error())
	}
	return nil
}

// liveDue takes the running fight's summary under the lock, so the
// upload above happens with nothing held.
func (p *Pipeline) liveDue(now time.Time) (string, fight.Fight, summary.Summary, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil || p.cur.ReportID == "" || now.Sub(p.lastLive) < p.o.LiveEvery {
		return "", fight.Fight{}, summary.Summary{}, false
	}
	f, sum, open := p.sess.Snapshot()
	if !open {
		return "", fight.Fight{}, summary.Summary{}, false
	}
	p.lastLive = now
	return p.cur.ReportID, f, sum, true
}
