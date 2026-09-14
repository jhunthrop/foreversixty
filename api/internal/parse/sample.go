package parse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"reflect"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/zstdx"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	logparquet "github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
)

// SampleSize is how many fights the raw-sample check re-parses, as the
// contract sets it: two, chosen at random.
const SampleSize = 2

// FlagTampered is the flag a report carries when its stored events do
// not match its raw bytes.
const FlagTampered = "tampered"

// Sample re-parses up to SampleSize of a report's fights from the raw
// chunks the companion uploaded and compares the events it gets with
// the events stored for those fights. A mismatch flags the report and
// withdraws its ranking rows.
//
// A report whose raw chunks have not all arrived is not evidence of
// anything: the check simply skips the fights it cannot cover.
func Sample(ctx context.Context, d Deps, reportID string) error {
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		return err
	}
	chunks, err := d.Reports.RawChunks(ctx, reportID)
	if err != nil {
		return err
	}
	if len(chunks) == 0 {
		return nil
	}
	fights, err := d.Reports.Fights(ctx, reportID)
	if err != nil {
		return err
	}

	// Everything that reads bytes back out of the bucket - the raw
	// chunks and the stored events - runs under its own deadline, the
	// same way Report's whole-file read does in parse.go. The terminal
	// Flag/RemoveReport writes in tampered stay on ctx, outside this
	// deadline, so a check that runs out of read time still records
	// what it found.
	readCtx, cancel := context.WithTimeout(ctx, SampleTimeout)
	defer cancel()

	for _, f := range pick(fights, SampleSize) {
		raw, ok, err := d.rawRange(readCtx, reportID, chunks, f)
		if err != nil {
			return err
		}
		if !ok {
			continue
		}
		if err := d.compareFight(ctx, readCtx, rep, f, raw); err != nil {
			return err
		}
	}
	return nil
}

// pick chooses up to n fights at random. Only fights with a recorded
// raw range can be checked at all.
func pick(fights []reports.FightEntry, n int) []reports.FightEntry {
	candidates := make([]reports.FightEntry, 0, len(fights))
	for _, f := range fights {
		if f.DurationMS > 0 {
			candidates = append(candidates, f)
		}
	}
	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	if len(candidates) > n {
		candidates = candidates[:n]
	}
	return candidates
}

// rawRange assembles the raw bytes covering a fight, and reports false
// when a chunk it needs has not been uploaded.
func (d Deps) rawRange(ctx context.Context, reportID string, chunks []reports.RawChunk, f reports.FightEntry) ([]byte, bool, error) {
	start, end, err := d.Reports.FightRawRange(ctx, reportID, f.Index)
	if err != nil || start == nil || end == nil {
		return nil, false, err
	}
	var (
		out    []byte
		cursor int64 = -1
	)
	for _, c := range chunks {
		if c.End <= *start || c.Start >= *end {
			continue
		}
		switch {
		case cursor < 0 && c.Start > *start:
			// The first chunk that overlaps the fight does not reach
			// back to where the fight itself starts: a gap at the head
			// of the range, nothing to compare.
			return nil, false, nil
		case cursor >= 0 && c.Start != cursor:
			// A gap in the middle of the fight: nothing to compare.
			return nil, false, nil
		}
		body, err := d.Objects.Get(ctx, reports.Keys(reportID).Raw(c.Start))
		if err != nil {
			return nil, false, nil
		}
		packed, err := io.ReadAll(io.LimitReader(body, zstdx.MaxChunk))
		body.Close()
		if err != nil {
			return nil, false, fmt.Errorf("parse: read raw chunk %s@%d: %w", reportID, c.Start, err)
		}
		decoded, err := zstdx.DecodeAll(packed, zstdx.MaxChunk)
		if err != nil {
			return nil, false, fmt.Errorf("parse: decode raw chunk %s@%d: %w", reportID, c.Start, err)
		}
		out = append(out, decoded...)
		cursor = c.End
	}
	if cursor < 0 || cursor < *end {
		return nil, false, nil
	}
	return out, true, nil
}

// compareFight re-parses raw and compares the fight it finds with the
// events stored for f. readCtx bounds the reads (the re-parse feed and
// the stored Parquet); ctx is the caller's own, unbounded by readCtx's
// deadline, and is what tampered's terminal writes run on, so a check
// that ran out of read time can still record what it found.
//
// The layout is pinned from the report's own health record: a fight in
// the middle of a night starts at a 4 MiB chunk boundary, which carries
// no header, so a session left to infer could read the lines with the
// wrong dialect and call an honest report tampered.
func (d Deps) compareFight(ctx, readCtx context.Context, rep reports.Report, f reports.FightEntry, raw []byte) error {
	s := session.New(engine.SessionOptionsFor(rep.ID, f.Start, true, layoutOf(rep)))
	var found *session.Closed
	if err := stream(readCtx, s, byteReader(raw), func(c session.Closed) error {
		if found == nil && c.Fight.Start.Equal(f.Start) {
			closed := c
			found = &closed
		}
		return nil
	}); err != nil {
		return err
	}
	if found == nil {
		// Inconclusive rather than damning: the range starts at a chunk
		// boundary, so its first line is usually a partial one, and a
		// re-parse that finds no fight is far more likely to be this
		// check's own blind spot than a forgery. Flagging here would
		// take an honest guild out of the rankings.
		d.logger().Info("parse", "op", "sample", "report", rep.ID, "fight", f.Index,
			"result", "no fight found in the raw range; not checked")
		return nil
	}
	stored, err := d.storedEvents(readCtx, rep.ID, f.Index)
	if err != nil {
		return err
	}
	if why := sameEvents(stored, found.Events); why != "" {
		return d.tampered(ctx, rep, f.Index, why)
	}
	return nil
}

// storedEvents reads a fight's Parquet back out of the bucket.
func (d Deps) storedEvents(ctx context.Context, reportID string, index int) ([]event.Event, error) {
	body, err := d.Objects.Get(ctx, reports.Keys(reportID).FightEvents(index))
	if err != nil {
		return nil, fmt.Errorf("parse: read events %s/%d: %w", reportID, index, err)
	}
	defer body.Close()
	b, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("parse: read events %s/%d: %w", reportID, index, err)
	}
	return logparquet.Unmarshal(b)
}

// layoutOf reads the layout a report's parse settled on.
func layoutOf(rep reports.Report) string {
	if len(rep.Health) == 0 {
		return ""
	}
	var h session.Health
	if err := json.Unmarshal(rep.Health, &h); err != nil {
		return ""
	}
	return h.Layout
}

// sameEvents compares two event streams as the Parquet schema sees
// them, which is the only part of an event that was ever stored, minus
// the two bookkeeping columns that cannot line up: the re-parse starts
// at a chunk boundary rather than at the top of the file, so its line
// numbers and byte offsets are its own. Everything that says what
// happened is compared.
func sameEvents(stored, reparsed []event.Event) string {
	if len(stored) != len(reparsed) {
		return fmt.Sprintf("%d events stored, the raw bytes give %d", len(stored), len(reparsed))
	}
	for i := range stored {
		a, b := comparableRow(stored[i]), comparableRow(reparsed[i])
		if !reflect.DeepEqual(a, b) {
			return fmt.Sprintf("event %d differs: stored %s, the raw bytes give %s", i, a.Event, b.Event)
		}
	}
	return ""
}

// comparableRow is an event as stored, with the line number and byte
// offset cleared.
func comparableRow(e event.Event) logparquet.Row {
	r := logparquet.RowOf(e)
	r.Line, r.Offset = 0, 0
	return r
}

// tampered flags a report and withdraws its ranking rows.
func (d Deps) tampered(ctx context.Context, rep reports.Report, index int, why string) error {
	d.logger().Warn("parse", "op", "sample", "report", rep.ID, "fight", index, "mismatch", why)
	if err := d.Reports.Flag(ctx, rep.ID, FlagTampered); err != nil {
		return err
	}
	if d.Rank != nil {
		return d.Rank.RemoveReport(ctx, rep.ID)
	}
	return nil
}

// Worker runs raw-sample checks out of band, so the companion's
// completion call returns at once.
type Worker struct {
	Deps  Deps
	queue chan string
	done  chan struct{}
	once  sync.Once

	// mu guards closed against a Schedule that races Close: Close takes
	// the write lock before it closes queue, and Schedule holds the
	// read lock for the whole of its send, so a Schedule already inside
	// its critical section always finishes before queue can close, and
	// one that has not started yet always sees closed true and skips
	// the channel entirely. Without this a Schedule racing Close is a
	// send on a closed channel - a panic, not a drop.
	mu     sync.RWMutex
	closed bool
}

// QueueDepth is how many reports may be waiting to be checked. Beyond
// it, a completion drops its check rather than blocking the request:
// the check is an audit, not a gate.
const QueueDepth = 256

// NewWorker returns a worker with an empty queue.
func NewWorker(d Deps) *Worker {
	return &Worker{Deps: d, queue: make(chan string, QueueDepth), done: make(chan struct{})}
}

// Schedule queues a report for checking. It never blocks, and a
// Schedule that loses the race with Close drops its report the same
// way one that finds the queue full does, rather than panicking.
func (w *Worker) Schedule(reportID string) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.closed {
		w.Deps.logger().Warn("parse", "op", "sample", "report", reportID, "err", "worker is closed")
		return
	}
	select {
	case w.queue <- reportID:
	default:
		w.Deps.logger().Warn("parse", "op", "sample", "report", reportID, "err", "queue is full")
	}
}

// Run consumes the queue until ctx is cancelled or Close is called. A
// report already queued when Close runs is still drained and checked:
// closing a channel does not discard what is already buffered in it.
func (w *Worker) Run(ctx context.Context) {
	defer close(w.done)
	for {
		select {
		case <-ctx.Done():
			return
		case id, ok := <-w.queue:
			if !ok {
				return
			}
			if err := Sample(ctx, w.Deps, id); err != nil {
				w.Deps.logger().Error("parse", "op", "sample", "report", id, "err", err)
			}
		}
	}
}

// SampleTimeout bounds the read-heavy part of one report's check: the
// raw chunk and stored-events reads. See Sample.
const SampleTimeout = 5 * time.Minute

// Close stops the worker and waits for the check in flight, if any, to
// finish. No Schedule call started after Close returns can reach the
// queue.
func (w *Worker) Close() {
	w.once.Do(func() {
		w.mu.Lock()
		w.closed = true
		close(w.queue)
		w.mu.Unlock()
	})
	<-w.done
}

// byteReader reads a byte slice, so the sample check can share stream.
func byteReader(b []byte) io.Reader { return &sliceReader{b: b} }

type sliceReader struct {
	b []byte
	i int
}

func (r *sliceReader) Read(p []byte) (int, error) {
	if r.i >= len(r.b) {
		return 0, io.EOF
	}
	n := copy(p, r.b[r.i:])
	r.i += n
	return n, nil
}
