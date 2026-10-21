// companion/internal/pipeline/drain.go
// Draining the queue. One item at a time, oldest first, stopping on
// the first failure worth retrying — which is what puts the fights on
// the server in the order the raid fought them.
package pipeline

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/logger"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

func marshalComplete(in client.Complete) ([]byte, error) { return json.Marshal(in) }

// MaxStateReads is how many drains may fail to read an item's report
// state before the item is dropped as poisoned. Network failures are
// not counted here and are never capped: an outage keeps the queue.
const MaxStateReads = 20

// Drain sends queued work until the queue is empty or an upload fails
// in a way that is worth retrying. The failing item stays at the head,
// so nothing behind it overtakes it.
func (p *Pipeline) Drain(ctx context.Context) error {
	// The open report is created as soon as the network allows, not
	// when its first fight closes, so the live snapshots of the very
	// first pull have somewhere to go. It is copied out from under
	// the lock first: create talks to the network, and Status must
	// not wait on that.
	if rep, ok := p.openCopy(); ok && !awaitingLogger(rep) {
		if err := p.settle(ctx, &rep); err != nil {
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

// awaitingLogger is true for an automatic report that has no id yet,
// has not seen its logger and has closed no fight. Creating it now
// would mean crediting it later by a patch; the logger's first line,
// which is also the first line of the first pull, is moments away. A
// report whose first fight closed without ever naming a logger is
// created anyway, and so is any report whose queue is being drained
// (send does not wait).
func awaitingLogger(rep state.Report) bool {
	return rep.ReportID == "" && rep.Auto && rep.DetectedName == "" && rep.Fights == 0
}

// openCopy copies the open report out from under the lock, so what
// settle does with the network does not hold it.
func (p *Pipeline) openCopy() (state.Report, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil {
		return state.Report{}, false
	}
	return *p.cur, true
}

// settle brings a report's server side up to date: it is created when
// it has no id, and its logging character is worked out and sent. The
// character goes with the creation when the combat log and the addon
// already name it, and by one PATCH later when the log named the
// logger before the addon's record of the character had arrived (the
// record is written when the player logs out). What changed is saved
// and handed back to the open report. A failure that waiting can fix
// is returned and nothing is lost: the next drain tries again.
func (p *Pipeline) settle(ctx context.Context, rep *state.Report) error {
	was := *rep
	p.resolve(rep)
	if rep.ReportID == "" {
		if err := p.create(ctx, rep); err != nil {
			return err
		}
	}
	if err := p.credit(ctx, rep); err != nil {
		return err
	}
	if rep.ReportID != was.ReportID || rep.Credited != was.Credited || rep.Auto != was.Auto ||
		rep.LoggingCharacter != was.LoggingCharacter {
		if err := state.Save(p.o.StateDir, *rep); err != nil {
			return err
		}
	}
	p.adopt(*rep)
	return nil
}

// resolve names the character an automatic report is credited to from
// what the combat log said and the addon has exported. It says so once
// per report when it cannot.
func (p *Pipeline) resolve(rep *state.Report) {
	if !rep.Auto || rep.LoggingCharacter != nil || rep.DetectedName == "" {
		return
	}
	var records []logger.Record
	if p.o.Records != nil {
		records = p.o.Records()
	}
	d := logger.Detected{Name: rep.DetectedName, Realm: rep.DetectedRealm, Class: rep.DetectedClass}
	c, why := logger.Judge(d, records)
	if why == logger.Matched {
		rep.LoggingCharacter = &c
		p.o.Log.Info("the log is a known character's", "component", "uploader",
			"report", rep.Key, "character", c.Key())
		return
	}
	p.mu.Lock()
	first := !p.noted[rep.Key]
	p.noted[rep.Key] = true
	p.mu.Unlock()
	if first {
		p.o.Log.Info(why.Unmatched(rep.DetectedName), "component", "uploader", "report", rep.Key)
	}
}

// credit sends the logging character of a report that already exists
// and was created without it. A refusal the server will repeat gives
// the attempt up: the report stays uncredited rather than being asked
// again on every tick.
func (p *Pipeline) credit(ctx context.Context, rep *state.Report) error {
	if !rep.Auto || rep.Credited || rep.LoggingCharacter == nil || rep.ReportID == "" {
		return nil
	}
	err := p.o.Client.UpdateReport(ctx, rep.ReportID, client.UpdateReport{LoggingCharacter: rep.LoggingCharacter})
	switch {
	case err == nil:
		rep.Credited = true
		p.o.Log.Info("a report was credited to its logger", "component", "uploader",
			"report", rep.Key, "report_id", rep.ReportID, "character", rep.LoggingCharacter.Key())
		return nil
	case client.Retryable(err) || client.Unauthorized(err):
		return err
	}
	p.o.Log.Error("the server refused the logging character; the report stays uncredited",
		"component", "uploader", "report", rep.Key, "character", rep.LoggingCharacter.Key(),
		"err", err.Error())
	rep.Auto = false
	return nil
}

// adopt copies what settle changed onto the open report, which Tick
// may have moved on since it was copied.
func (p *Pipeline) adopt(rep state.Report) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil || p.cur.Key != rep.Key {
		return
	}
	p.cur.ReportID, p.cur.LoggingCharacter = rep.ReportID, rep.LoggingCharacter
	p.cur.Credited, p.cur.Auto = rep.Credited, rep.Auto
	if name := detectedLabel(*p.cur); name != "" {
		p.lastDetected = name
	}
}

// create asks the API for a report id and records it on rep.
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
	rep.Credited = rep.LoggingCharacter != nil
	p.o.Log.Info("a report was created", "component", "uploader",
		"report", rep.Key, "report_id", rep.ReportID)
	return nil
}

// send performs one queued item.
func (p *Pipeline) send(ctx context.Context, l *queue.Lease) error {
	rep, err := state.Load(p.o.StateDir, l.Item.ReportKey)
	if errors.Is(err, fs.ErrNotExist) {
		p.o.Log.Error("dropping a queued item whose report state is gone",
			"component", "uploader", "report", l.Item.ReportKey,
			"kind", string(l.Item.Kind), "err", err.Error())
		return l.Drop()
	}
	if err != nil {
		// The file is there and could not be read: a full descriptor
		// table, a disk error, a half-written file. That is the same
		// shape as a retryable upload failure, so the item stays at
		// the head and the drain stops rather than throwing a fight
		// away over a condition that passes. A file that never
		// becomes readable is a different thing: after MaxStateReads
		// drains it is treated as corrupt and the item is dropped, so
		// one bad byte cannot hold every later night's uploads.
		if l.Item.Attempts+1 >= MaxStateReads {
			p.o.Log.Error("dropping a queued item whose report state never became readable",
				"component", "uploader", "report", l.Item.ReportKey,
				"kind", string(l.Item.Kind), "attempts", l.Item.Attempts+1, "err", err.Error())
			return l.Drop()
		}
		if ferr := l.Fail(err); ferr != nil {
			return ferr
		}
		return fmt.Errorf("read the state of report %s: %w", l.Item.ReportKey, err)
	}
	if serr := p.settle(ctx, &rep); serr != nil {
		return p.failed(l, serr)
	}

	// stored stays Created for anything the API does not answer
	// idempotently, which is everything but the two writes below.
	var stored client.Stored
	switch l.Item.Kind {
	case queue.Fight:
		stored, err = p.o.Client.PutFight(ctx, rep.ReportID, l.Item.FightIndex,
			l.Item.ContentType, l.Body)
	case queue.Raw:
		stored, err = p.o.Client.PutRaw(ctx, rep.ReportID, l.Item.Offset, l.Item.SHA256, l.Body)
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
	if stored == client.Duplicate {
		// At-least-once delivery is safe by contract, so this is not
		// a problem — but on a resume it is the line that says how
		// much of the night the server already had.
		p.o.Log.Debug("the server already held this item", "component", "uploader",
			"report", rep.Key, "report_id", rep.ReportID, "kind", string(l.Item.Kind),
			"fight", l.Item.FightIndex, "offset", l.Item.Offset)
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
	// Queued is every item waiting (fights, raw chunks, the completion);
	// QueuedFights is the fights among them, the ones the window counts.
	// Fights is every fight the report has closed, so the fights the
	// site holds are Fights minus QueuedFights.
	Queued       int `json:"queued"`
	QueuedFights int `json:"queued_fights"`
	// Detected is who the combat log says is logging: the report's
	// character once the addon has named it, else the name the log
	// gave. It is the last one seen when no report is open.
	Detected string `json:"detected,omitempty"`
}

// Status reports what the pipeline is doing right now.
func (p *Pipeline) Status() Status {
	s := Status{}
	if n, err := p.o.Queue.Len(); err == nil {
		s.Queued = n
	}
	if n, err := p.o.Queue.Waiting(queue.Fight); err == nil {
		s.QueuedFights = n
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	s.Detected = p.lastDetected
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
