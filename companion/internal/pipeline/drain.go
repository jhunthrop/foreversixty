// companion/internal/pipeline/drain.go
// Draining the queue. One item at a time, oldest first, stopping on
// the first failure worth retrying — which is what puts the fights on
// the server in the order the raid fought them.
package pipeline

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

func marshalComplete(in client.Complete) ([]byte, error) { return json.Marshal(in) }

// Drain sends queued work until the queue is empty or an upload fails
// in a way that is worth retrying. The failing item stays at the head,
// so nothing behind it overtakes it.
func (p *Pipeline) Drain(ctx context.Context) error {
	// The open report is created as soon as the network allows, not
	// when its first fight closes, so the live snapshots of the very
	// first pull have somewhere to go. It is copied out from under
	// the lock first: create talks to the network, and Status must
	// not wait on that.
	if rep, ok := p.openWithoutID(); ok {
		if err := p.create(ctx, &rep); err != nil {
			return err
		}
		if err := state.Save(p.o.StateDir, rep); err != nil {
			return err
		}
	}
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		l, err := p.o.Queue.Head()
		if errors.Is(err, queue.ErrEmpty) {
			return nil
		}
		if err != nil {
			return err
		}
		if err := p.send(ctx, l); err != nil {
			return err
		}
	}
}

// openWithoutID copies the open report when it has no server id yet.
func (p *Pipeline) openWithoutID() (state.Report, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil || p.cur.ReportID != "" {
		return state.Report{}, false
	}
	return *p.cur, true
}

// create asks the API for a report id and records it on rep and, when
// rep is the open report, on the pipeline.
func (p *Pipeline) create(ctx context.Context, rep *state.Report) error {
	created, err := p.o.Client.CreateReport(ctx, client.CreateReport{
		Title:      rep.Title,
		Visibility: rep.Visibility,
		Zone:       rep.Zone,
		// The attribution triple carries the ruleset segment through
		// untouched; the API maps it.
		LoggingCharacter: rep.LoggingCharacter,
	})
	if err != nil {
		return err
	}
	rep.ReportID = created.ID
	p.mu.Lock()
	if p.cur != nil && p.cur.Key == rep.Key {
		p.cur.ReportID = created.ID
	}
	p.mu.Unlock()
	p.o.Log.Info("a report was created", "component", "uploader",
		"report", rep.Key, "report_id", rep.ReportID)
	return nil
}

// send performs one queued item.
func (p *Pipeline) send(ctx context.Context, l *queue.Lease) error {
	rep, err := state.Load(p.o.StateDir, l.Item.ReportKey)
	if err != nil {
		p.o.Log.Error("dropping a queued item whose report state is gone",
			"component", "uploader", "report", l.Item.ReportKey,
			"kind", string(l.Item.Kind), "err", err.Error())
		return l.Drop()
	}
	if rep.ReportID == "" {
		if cerr := p.create(ctx, &rep); cerr != nil {
			return p.failed(l, cerr)
		}
		if err := state.Save(p.o.StateDir, rep); err != nil {
			return err
		}
	}

	switch l.Item.Kind {
	case queue.Fight:
		_, err = p.o.Client.PutFight(ctx, rep.ReportID, l.Item.FightIndex,
			l.Item.ContentType, l.Body)
	case queue.Raw:
		_, err = p.o.Client.PutRaw(ctx, rep.ReportID, l.Item.Offset, l.Item.SHA256, l.Body)
	case queue.Complete:
		var in client.Complete
		if jerr := json.Unmarshal(l.Body, &in); jerr != nil {
			p.o.Log.Error("dropping an unreadable completion", "component", "uploader",
				"report", rep.Key, "err", jerr.Error())
			return l.Drop()
		}
		err = p.o.Client.Complete(ctx, rep.ReportID, in)
	default:
		p.o.Log.Error("dropping a queued item of an unknown kind",
			"component", "uploader", "kind", string(l.Item.Kind))
		return l.Drop()
	}
	if err != nil {
		return p.failed(l, err)
	}
	if l.Item.Kind == queue.Complete {
		rep.Done = true
		if err := state.Save(p.o.StateDir, rep); err != nil {
			return err
		}
	}
	return l.Ack()
}

// failed decides what a failed upload means. Anything retryable leaves
// the item where it is and stops the drain; anything the server will
// refuse again forever is dropped with a loud log line, because a
// poisoned item must not hold a raid night hostage.
func (p *Pipeline) failed(l *queue.Lease, cause error) error {
	if client.Retryable(cause) || client.Unauthorized(cause) {
		if err := l.Fail(cause); err != nil {
			return err
		}
		return cause
	}
	p.o.Log.Error("dropping an item the server refused", "component", "uploader",
		"report", l.Item.ReportKey, "kind", string(l.Item.Kind),
		"fight", l.Item.FightIndex, "err", cause.Error())
	return l.Drop()
}

// Status is what the tray and the status page show.
type Status struct {
	Logging    bool      `json:"logging"`
	ReportKey  string    `json:"report_key,omitempty"`
	ReportID   string    `json:"report_id,omitempty"`
	LogPath    string    `json:"log_path,omitempty"`
	Fights     int       `json:"fights"`
	Offset     int64     `json:"offset"`
	LastAppend time.Time `json:"last_append,omitzero"`
	Queued     int       `json:"queued"`
}

// Status reports what the pipeline is doing right now.
func (p *Pipeline) Status() Status {
	s := Status{}
	if n, err := p.o.Queue.Len(); err == nil {
		s.Queued = n
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil {
		return s
	}
	s.Logging = true
	s.ReportKey, s.ReportID = p.cur.Key, p.cur.ReportID
	s.LogPath, s.Fights = p.cur.LogPath, p.cur.Fights
	s.Offset, s.LastAppend = p.cur.Offset, p.cur.LastAppend
	return s
}

// Run drives the pipeline until the context ends: parse on every tick,
// upload what is queued, and push a live snapshot when one is due.
// Errors are logged rather than returned, because a companion that
// exits on a failed upload is a companion that loses a raid night.
func (p *Pipeline) Run(ctx context.Context, every time.Duration) error {
	t := time.NewTicker(every)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case now := <-t.C:
			if err := p.Tick(now); err != nil {
				p.o.Log.Error("the tail failed", "component", "pipeline", "err", err.Error())
			}
			if err := p.Live(ctx, now); err != nil {
				p.o.Log.Error("a live snapshot failed", "component", "pipeline", "err", err.Error())
			}
			if err := p.Drain(ctx); err != nil && !errors.Is(err, context.Canceled) {
				p.o.Log.Warn("uploads are behind", "component", "uploader", "err", err.Error())
			}
		}
	}
}
