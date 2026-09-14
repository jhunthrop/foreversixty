package pipeline_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
	"github.com/jhunthrop/foreversixty/companion/internal/pipeline"
	"github.com/jhunthrop/foreversixty/companion/internal/queue"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
	"github.com/jhunthrop/foreversixty/companion/internal/watch"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// rig is one companion wired to one fake API over one temporary game
// directory.
type rig struct {
	t        *testing.T
	srv      *fakeapi.Server
	dirs     struct{ logs, queue, state string }
	log      string
	pipe     *pipeline.Pipeline
	q        *queue.Queue
	clock    time.Time
	rawChunk int
}

func newRig(t *testing.T, rawChunk int) *rig {
	t.Helper()
	root := t.TempDir()
	r := &rig{t: t, srv: fakeapi.New(), clock: t0, rawChunk: rawChunk}
	r.dirs.logs = filepath.Join(root, "Logs")
	r.dirs.queue = filepath.Join(root, "queue")
	r.dirs.state = filepath.Join(root, "state")
	for _, d := range []string{r.dirs.logs, r.dirs.queue, r.dirs.state} {
		if err := os.MkdirAll(d, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	r.log = filepath.Join(r.dirs.logs, "WoWCombatLog.txt")
	t.Cleanup(r.srv.Close)
	r.open()
	// The companion starts before the player logs in: one tick over
	// an empty Logs directory, so the log it later sees is a new one
	// rather than a night it never watched.
	if err := r.pipe.Tick(r.clock); err != nil {
		t.Fatal(err)
	}
	return r
}

// open builds a fresh pipeline over the same directories, which is
// what a restart looks like.
func (r *rig) open() {
	r.t.Helper()
	if r.pipe != nil {
		r.pipe.Close()
	}
	c, err := client.New(client.Options{
		BaseURL: r.srv.URL,
		Token:   func() string { return fakeapi.Token },
		Retry:   client.Retry{MaxAttempts: 1, Base: time.Millisecond, Max: time.Millisecond},
	})
	if err != nil {
		r.t.Fatal(err)
	}
	q, err := queue.Open(r.dirs.queue)
	if err != nil {
		r.t.Fatal(err)
	}
	r.q = q
	p, err := pipeline.New(pipeline.Options{
		StateDir: r.dirs.state,
		Client:   c,
		Queue:    q,
		Watch:    watch.New(watch.Options{Dir: r.dirs.logs}),
		Config:   config.Config{APIBaseURL: r.srv.URL, SiteBaseURL: r.srv.URL, ReportVisibility: "public"},
		RawChunk: r.rawChunk,
		// Save the session on every tick so a restart in a test
		// resumes from the same place a long-running companion would.
		StateEvery: -1,
	})
	if err != nil {
		r.t.Fatal(err)
	}
	r.pipe = p
	r.t.Cleanup(func() { p.Close() })
}

// write appends to the game's log and advances the clock.
func (r *rig) write(text string) {
	r.t.Helper()
	f, err := os.OpenFile(r.log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		r.t.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		r.t.Fatal(err)
	}
	f.Close()
	r.clock = r.clock.Add(time.Second)
	if err := os.Chtimes(r.log, r.clock, r.clock); err != nil {
		r.t.Fatal(err)
	}
}

// tick parses and uploads once.
func (r *rig) tick() error {
	r.t.Helper()
	r.clock = r.clock.Add(time.Second)
	if err := r.pipe.Tick(r.clock); err != nil {
		return err
	}
	return r.pipe.Drain(r.t.Context())
}

// settle ticks until the log has been quiet long enough to complete
// the report and the queue has drained.
func (r *rig) settle() {
	r.t.Helper()
	if err := r.tick(); err != nil {
		r.t.Fatalf("tick: %v", err)
	}
	r.clock = r.clock.Add(watch.CompleteIdle + time.Minute)
	if err := r.pipe.Tick(r.clock); err != nil {
		r.t.Fatalf("tick: %v", err)
	}
	if err := r.pipe.Drain(r.t.Context()); err != nil {
		r.t.Fatalf("drain: %v", err)
	}
}

func TestAWholeNightArrivesAsFightsInOrder(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	for i := range 3 {
		r.write(fixture.Encounter(i) + fixture.Heartbeat(i))
		if err := r.tick(); err != nil {
			t.Fatal(err)
		}
	}
	r.settle()

	var fights []string
	var completed bool
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights = append(fights, e)
		}
		if strings.HasPrefix(e, "complete ") {
			completed = true
		}
	}
	if len(fights) < 3 {
		t.Fatalf("the server saw %v", r.srv.Order())
	}
	for i := 1; i < len(fights); i++ {
		if fights[i-1] >= fights[i] {
			t.Fatalf("fights arrived out of order: %v", fights)
		}
	}
	if !completed {
		t.Errorf("the report was never completed: %v", r.srv.Order())
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Errorf("%d items are still queued", n)
	}
	for _, rep := range r.srv.Reports() {
		if len(rep.Raw) == 0 {
			t.Error("no raw chunks were uploaded")
		}
		for _, f := range rep.Fights {
			if len(f.Events) == 0 {
				t.Errorf("fight %d has no parquet", f.Index)
			}
			if len(f.Metrics) == 0 {
				t.Errorf("fight %d has no metrics rows", f.Index)
			}
			if len(f.RawRange.SHA256) != 64 {
				t.Errorf("fight %d raw hash = %q", f.Index, f.RawRange.SHA256)
			}
		}
	}
}

func TestRawIsChunkedAtTheConfiguredSizeAndAddressedByOffset(t *testing.T) {
	r := newRig(t, 512)
	r.write(fixture.Log(2))
	for range 4 {
		if err := r.tick(); err != nil {
			t.Fatal(err)
		}
	}
	r.settle()
	for _, rep := range r.srv.Reports() {
		if len(rep.Raw) < 2 {
			t.Fatalf("raw chunks = %d, want several at 512 bytes", len(rep.Raw))
		}
		for offset, body := range rep.Raw {
			if offset%512 != 0 {
				t.Errorf("chunk at offset %d is not on a 512-byte boundary", offset)
			}
			if len(body) == 0 {
				t.Errorf("chunk at %d is empty", offset)
			}
		}
		if _, first := rep.Raw[0]; !first {
			t.Error("there is no chunk at offset zero")
		}
	}
}

func TestTheRawHashIsOverTheDecodedChunk(t *testing.T) {
	// The fake ingest decodes each chunk and compares the header
	// against the decoded bytes, so a digest taken over the
	// compressed frame would fail the upload outright.
	r := newRig(t, 512)
	r.write(fixture.Log(1))
	for range 3 {
		if err := r.tick(); err != nil {
			t.Fatal(err)
		}
	}
	r.settle()
	for _, rep := range r.srv.Reports() {
		if len(rep.Raw) == 0 {
			t.Fatal("no raw chunks were accepted")
		}
	}
}

func TestAnOutageQueuesAndTheOrderSurvivesTheReconnect(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}

	r.srv.Offline(true)
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	r.write(fixture.Encounter(2) + fixture.Heartbeat(2))
	if err := r.tick(); err == nil {
		t.Fatal("draining while offline succeeded")
	}
	if n, _ := r.q.Len(); n == 0 {
		t.Fatal("nothing was queued during the outage")
	}

	r.srv.Offline(false)
	r.settle()
	var fights []string
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights = append(fights, e)
		}
	}
	if len(fights) < 3 {
		t.Fatalf("after the reconnect the server saw %v", r.srv.Order())
	}
	for i := 1; i < len(fights); i++ {
		if fights[i-1] >= fights[i] {
			t.Fatalf("fights arrived out of order: %v", fights)
		}
	}
}

func TestARestartResumesTheSameReport(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	before, err := state.List(r.dirs.state)
	if err != nil || len(before) != 1 {
		t.Fatalf("state = %+v, %v", before, err)
	}
	key, reportID := before[0].Key, before[0].ReportID

	r.open() // the companion restarts
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.settle()

	after, err := state.List(r.dirs.state)
	if err != nil {
		t.Fatal(err)
	}
	if len(after) != 1 {
		t.Fatalf("the restart started a second report: %+v", after)
	}
	if after[0].Key != key {
		t.Fatalf("report key = %q, want %q", after[0].Key, key)
	}
	if len(r.srv.Reports()) != 1 {
		t.Fatalf("the server holds %d reports", len(r.srv.Reports()))
	}
	if r.srv.Reports()[reportID] == nil {
		t.Fatalf("report %q is gone", reportID)
	}
	if got := len(r.srv.Reports()[reportID].Fights); got < 2 {
		t.Errorf("the resumed report holds %d fights, want both", got)
	}
}

func TestALiveSnapshotGoesUpWhileAFightIsOpen(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	// An encounter with no end: the fight stays open.
	open := strings.SplitAfter(fixture.Encounter(0), "\n")
	r.write(strings.Join(open[:len(open)-2], ""))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.clock = r.clock.Add(10 * time.Second)
	if err := r.pipe.Live(t.Context(), r.clock); err != nil {
		t.Fatal(err)
	}
	var live int
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "live ") {
			live++
		}
	}
	if live != 1 {
		t.Fatalf("live snapshots = %d, order = %v", live, r.srv.Order())
	}
	st := r.pipe.Status()
	if !st.Logging || st.ReportID == "" {
		t.Fatalf("status = %+v", st)
	}
}

func TestNothingHappensWithoutALogFile(t *testing.T) {
	r := newRig(t, 1<<20)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	if len(r.srv.Order()) != 0 {
		t.Fatalf("the server saw %v with no log file", r.srv.Order())
	}
	if st := r.pipe.Status(); st.Logging {
		t.Errorf("status = %+v", st)
	}
}

func TestRunTicksUntilTheContextEnds(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan error, 1)
	go func() { done <- r.pipe.Run(ctx, time.Millisecond) }()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if len(r.srv.Order()) > 0 {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatalf("Run = %v", err)
	}
	if len(r.srv.Order()) == 0 {
		t.Fatal("Run uploaded nothing")
	}
}

func TestAFightTheServerRefusesIsDroppedSoTheQueueMoves(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	// A verification mismatch: permanent, so the fight is dropped
	// with a log line rather than retried until the heat death.
	r.srv.RefuseFights(true)
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	if err := r.tick(); err != nil {
		t.Fatalf("a refused fight stopped the drain: %v", err)
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Fatalf("%d items are stuck behind the refused fight", n)
	}
	r.srv.RefuseFights(false)
	r.write(fixture.Encounter(2) + fixture.Heartbeat(2))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	var fights int
	for _, e := range r.srv.Order() {
		if strings.HasPrefix(e, "fight ") {
			fights++
		}
	}
	if fights != 2 {
		t.Fatalf("the server stored %d fights, want the two it accepted: %v", fights, r.srv.Order())
	}
}

func TestAQueuedItemWithNoReportStateIsDropped(t *testing.T) {
	r := newRig(t, 1<<20)
	if _, err := r.q.Enqueue(queue.Item{
		Kind: queue.Fight, ReportKey: "a-report-that-never-existed", ContentType: "text/plain",
	}, []byte("x")); err != nil {
		t.Fatal(err)
	}
	if err := r.pipe.Drain(t.Context()); err != nil {
		t.Fatal(err)
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Fatalf("%d orphans are still queued", n)
	}
}

func TestLiveAndCompleteDoNothingWithoutAnOpenReport(t *testing.T) {
	r := newRig(t, 1<<20)
	if err := r.pipe.Live(t.Context(), r.clock); err != nil {
		t.Fatal(err)
	}
	// A completion with nothing open is what the watcher emits after
	// the companion starts on a quiet machine.
	r.clock = r.clock.Add(watch.CompleteIdle + time.Minute)
	if err := r.pipe.Tick(r.clock); err != nil {
		t.Fatal(err)
	}
	if len(r.srv.Order()) != 0 {
		t.Fatalf("the server saw %v", r.srv.Order())
	}
}

// unpack decodes a raw chunk the way the ingest route does, so a test
// can compare what the server would reassemble against the log.
func unpack(t *testing.T, packed []byte) []byte {
	t.Helper()
	d, err := zstd.NewReader(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	out, err := d.DecodeAll(packed, nil)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func TestASecondReportInTheSameFileIsAddressedFromItsOwnFirstByte(t *testing.T) {
	// Every other test starts its report at file offset zero, where a
	// file offset and a report offset are the same number. Here the
	// thirty-minute rule opens a second report partway through one log
	// file, so the two differ and the contract's "all offsets are
	// report-relative" is the only thing that can hold.
	r := newRig(t, 1<<20)
	first := fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0)
	r.write(first)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.clock = r.clock.Add(watch.NewReportIdle + time.Minute)
	rest := fixture.Encounter(1) + fixture.Heartbeat(1)
	r.write(rest)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.settle()

	reports, err := state.List(r.dirs.state)
	if err != nil {
		t.Fatal(err)
	}
	if len(reports) != 2 {
		t.Fatalf("state holds %d reports, want two: %+v", len(reports), reports)
	}
	next := reports[1]
	if next.StartOffset != int64(len(first)) {
		t.Fatalf("the second report starts at file offset %d, want %d",
			next.StartOffset, len(first))
	}
	rep := r.srv.Reports()[next.ReportID]
	if rep == nil {
		t.Fatalf("the server does not hold %q: %v", next.ReportID, r.srv.Order())
	}
	// Addressed from zero, not from 1500: a file-offset implementation
	// leaves a hole at the front of the second report's raw stream.
	wantRawStream(t, rep, rest)
	if rep.Complete == nil {
		t.Fatal("the second report was never completed")
	}
	if rep.Complete.FinalOffset != int64(len(rest)) {
		t.Errorf("final_offset = %d, want %d — the file offset would be %d",
			rep.Complete.FinalOffset, len(rest), len(first)+len(rest))
	}
	for _, f := range rep.Fights {
		if f.RawRange.StartOffset != 0 {
			t.Errorf("fight %d raw_range.start_offset = %d, want 0 — the file offset would be %d",
				f.Index, f.RawRange.StartOffset, next.StartOffset)
		}
	}
}

// wantRawStream reassembles the report's raw chunks the way the server
// would and checks they are exactly the bytes the engine parsed. A
// gap means bytes were never uploaded; a length mismatch means bytes
// went up twice.
func wantRawStream(t *testing.T, rep *fakeapi.Report, want string) {
	t.Helper()
	var whole []byte
	for _, offset := range rawOffsets(rep) {
		if offset != int64(len(whole)) {
			t.Fatalf("the raw stream jumps from %d to %d: chunks at %v",
				len(whole), offset, rawOffsets(rep))
		}
		whole = append(whole, unpack(t, rep.Raw[offset])...)
	}
	if string(whole) == want {
		return
	}
	if len(whole) != len(want) {
		t.Errorf("the reassembled raw stream is %d bytes, want the %d the engine parsed",
			len(whole), len(want))
		return
	}
	for i := range whole {
		if whole[i] != want[i] {
			t.Errorf("the reassembled raw stream differs from the log at byte %d", i)
			return
		}
	}
}

func rawOffsets(rep *fakeapi.Report) []int64 {
	offsets := make([]int64, 0, len(rep.Raw))
	for o := range rep.Raw {
		offsets = append(offsets, o)
	}
	slices.Sort(offsets)
	return offsets
}

func TestARestartUploadsTheRawBytesThatWereStillBuffered(t *testing.T) {
	// The bytes between the last queued chunk and the stopping point
	// live only in memory, so a restart has to read them back out of
	// the log; without that the server holds a raw stream with a hole
	// at the front and can never reassemble it.
	r := newRig(t, 1<<20) // far larger than the fixture: nothing flushes until completion
	first := fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0)
	r.write(first)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}

	r.open() // the companion restarts with those bytes unqueued
	rest := fixture.Encounter(1) + fixture.Heartbeat(1)
	r.write(rest)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.settle()

	reports := r.srv.Reports()
	if len(reports) != 1 {
		t.Fatalf("the server holds %d reports, want one", len(reports))
	}
	for _, rep := range reports {
		wantRawStream(t, rep, first+rest)
	}
}

func TestARestartBeforeAnyLayoutSettlesRefeedsTheReport(t *testing.T) {
	// A report that begins partway through a file has no header, so
	// until inference settles a layout the engine refuses to serialise
	// the session. Nothing has closed, so there is nothing to lose:
	// the restart re-feeds the report from its first byte, and the
	// bytes it re-reads must not be uploaded a second time.
	r := newRig(t, 1<<20)
	opening := fixture.Zone
	r.write(opening)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	saved, err := state.List(r.dirs.state)
	if err != nil || len(saved) != 1 {
		t.Fatalf("state = %+v, %v", saved, err)
	}
	if len(saved[0].Session) != 0 {
		t.Fatalf("a session with no settled layout was serialised anyway (%d bytes)",
			len(saved[0].Session))
	}
	key := saved[0].Key

	r.open() // the companion restarts with no session blob to restore
	rest := fixture.Encounter(0) + fixture.Heartbeat(0)
	r.write(rest)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.settle()

	after, err := state.List(r.dirs.state)
	if err != nil {
		t.Fatal(err)
	}
	if len(after) != 1 || after[0].Key != key {
		t.Fatalf("the restart did not continue report %q: %+v", key, after)
	}
	reports := r.srv.Reports()
	if len(reports) != 1 {
		t.Fatalf("the server holds %d reports, want one", len(reports))
	}
	for _, rep := range reports {
		// The re-fed bytes are the point: they must appear once.
		wantRawStream(t, rep, opening+rest)
	}
}

func TestAReportWhoseLogVanishedIsAbandonedAndTheNextOneStarts(t *testing.T) {
	// The player moved the install, or emptied the Logs folder. The
	// report that was open cannot be continued, so it is closed out
	// and the companion goes on logging rather than failing to start.
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	before, err := state.List(r.dirs.state)
	if err != nil || len(before) != 1 {
		t.Fatalf("state = %+v, %v", before, err)
	}
	gone := before[0].Key

	if err := os.Remove(r.log); err != nil {
		t.Fatal(err)
	}
	r.open() // the companion restarts with the log missing

	if st := r.pipe.Status(); st.Logging {
		t.Errorf("the companion resumed a report whose log is gone: %+v", st)
	}
	abandoned, err := state.Load(r.dirs.state, gone)
	if err != nil {
		t.Fatal(err)
	}
	if !abandoned.Closed {
		t.Errorf("report %q is still open, so every restart would retry it", gone)
	}

	// A new night in a new file: the companion must pick it up.
	if err := r.pipe.Tick(r.clock); err != nil {
		t.Fatal(err)
	}
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(1) + fixture.Heartbeat(1))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.settle()

	after, err := state.List(r.dirs.state)
	if err != nil {
		t.Fatal(err)
	}
	if len(after) != 2 {
		t.Fatalf("state holds %d reports, want the abandoned one and the new one: %+v", len(after), after)
	}
	fresh := after[1]
	if fresh.Key == gone {
		t.Fatalf("the new night reused the abandoned report's key %q", gone)
	}
	rep := r.srv.Reports()[fresh.ReportID]
	if rep == nil {
		t.Fatalf("the server does not hold the new report %q: %v", fresh.ReportID, r.srv.Order())
	}
	if len(rep.Fights) != 1 {
		t.Errorf("the new report holds %d fights, want the one it logged", len(rep.Fights))
	}
}
