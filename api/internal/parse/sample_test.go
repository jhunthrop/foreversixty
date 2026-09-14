package parse

import (
	"context"
	"crypto/sha256"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	logparquet "github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

func TestSampleAcceptsAReportWhoseRawMatches(t *testing.T) {
	d, reportID, objects, rank := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	seedRawChunk(t, d, objects, reportID, engine.FixtureLog())

	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged != nil {
		t.Fatalf("an honest report was flagged %q", *rep.Flagged)
	}
	if len(rank.removed) != 0 {
		t.Fatalf("rankings were withdrawn from an honest report: %v", rank.removed)
	}
}

func TestSampleFlagsAReportWhoseStoredEventsDoNotMatch(t *testing.T) {
	d, reportID, objects, rank := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	seedRawChunk(t, d, objects, reportID, engine.FixtureLog())

	trimStoredEvents(t, d, objects, reportID)

	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged == nil || *rep.Flagged != FlagTampered {
		t.Fatalf("flagged = %v, want tampered", rep.Flagged)
	}
	if len(rank.removed) != 1 || rank.removed[0] != reportID {
		t.Fatalf("removed = %v, want the report", rank.removed)
	}
}

func TestSampleSkipsAReportWithNoRawChunks(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged != nil {
		t.Fatalf("a report with nothing to check was flagged %q", *rep.Flagged)
	}
}

// seedRawChunk stores one raw chunk covering the whole log, as the
// companion would have uploaded it: report-relative offset zero.
func seedRawChunk(t *testing.T, d Deps, objects *memObjects, reportID string, raw []byte) {
	t.Helper()
	ctx := context.Background()
	sum := sha256.Sum256(raw)
	if _, err := d.Reports.PutRawChunk(ctx, reportID, reports.RawChunk{
		Start: 0, End: int64(len(raw)), SHA256: sum[:],
	}); err != nil {
		t.Fatal(err)
	}
	if err := objects.Put(ctx, store.Keys{ReportID: reportID}.Raw(0), pack(t, raw), store.PutOptions{}); err != nil {
		t.Fatal(err)
	}
}

func TestTheWorkerRunsAScheduledCheckAndStops(t *testing.T) {
	d, reportID, objects, _ := fixtureUpload(t, false)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	seedRawChunk(t, d, objects, reportID, engine.FixtureLog())

	w := NewWorker(d)
	done := make(chan struct{})
	go func() { w.Run(ctx); close(done) }()
	w.Schedule(reportID)
	w.Close()
	<-done
}

func TestSchedulingBeyondTheQueueDepthDropsRatherThanBlocks(t *testing.T) {
	d, _, _, _ := fixtureUpload(t, false)
	w := NewWorker(d)
	for range QueueDepth + 10 {
		w.Schedule("report")
	}
}

func TestPickChoosesAtMostTwoFights(t *testing.T) {
	fights := []reports.FightEntry{}
	for i := range 5 {
		var e reports.FightEntry
		e.Index, e.DurationMS = i, 1000
		fights = append(fights, e)
	}
	if got := pick(fights, SampleSize); len(got) != SampleSize {
		t.Fatalf("picked %d, want %d", len(got), SampleSize)
	}
	if got := pick(nil, SampleSize); len(got) != 0 {
		t.Fatalf("picked %d from nothing", len(got))
	}
}

func TestSampleSkipsAFightWhoseChunkIsMissing(t *testing.T) {
	d, reportID, _, rank := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	// A chunk is recorded but its object never arrived.
	raw := engine.FixtureLog()
	sum := sha256.Sum256(raw)
	if _, err := d.Reports.PutRawChunk(ctx, reportID, reports.RawChunk{
		Start: 0, End: int64(len(raw)), SHA256: sum[:],
	}); err != nil {
		t.Fatal(err)
	}
	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged != nil {
		t.Fatalf("a report with an unuploaded chunk was flagged %q", *rep.Flagged)
	}
	if len(rank.removed) != 0 {
		t.Fatal("nothing should have been withdrawn")
	}
}

func TestSampleSkipsWhenTheRangeIsNotCovered(t *testing.T) {
	d, reportID, objects, _ := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	// Only the first half of the log was uploaded.
	raw := engine.FixtureLog()
	half := raw[:len(raw)/2]
	sum := sha256.Sum256(half)
	if _, err := d.Reports.PutRawChunk(ctx, reportID, reports.RawChunk{
		Start: 0, End: int64(len(half)), SHA256: sum[:],
	}); err != nil {
		t.Fatal(err)
	}
	if err := objects.Put(ctx, store.Keys{ReportID: reportID}.Raw(0), pack(t, half), store.PutOptions{}); err != nil {
		t.Fatal(err)
	}
	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged != nil {
		t.Fatalf("a half-uploaded report was flagged %q", *rep.Flagged)
	}
}

func TestSampleSkipsAFightWhoseRangeHasAGapAtItsHead(t *testing.T) {
	d, reportID, objects, rank := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	raw := engine.FixtureLog()
	// Only the back half of the log arrived, but the fight's own
	// recorded start is forced to the very beginning of the file: the
	// one chunk that arrived does not reach back far enough to cover
	// it, which is a gap at the head of the range, not the middle.
	half := int64(len(raw) / 2)
	tail := raw[half:]
	sum := sha256.Sum256(tail)
	if _, err := d.Reports.PutRawChunk(ctx, reportID, reports.RawChunk{
		Start: half, End: int64(len(raw)), SHA256: sum[:],
	}); err != nil {
		t.Fatal(err)
	}
	if err := objects.Put(ctx, store.Keys{ReportID: reportID}.Raw(half), pack(t, tail), store.PutOptions{}); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Reports.Pool.Exec(ctx,
		`update fights set raw_start_offset = 0 where report_id = $1`, reportID); err != nil {
		t.Fatal(err)
	}

	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged != nil {
		t.Fatalf("a report missing the head of its range was flagged %q", *rep.Flagged)
	}
	if len(rank.removed) != 0 {
		t.Fatal("nothing should have been withdrawn")
	}
}

// TestRawRangeRejectsAGapAtTheHeadOfTheFightsRange exercises rawRange
// directly rather than through Sample: compareFight's found == nil
// fallback would mask this bug too (a fight whose opening bytes are
// missing almost never re-parses to a fight matching f.Start either),
// so a Sample-level assertion on rep.Flagged cannot tell a fixed
// rawRange from a broken one. Only checking rawRange's own ok return
// proves the coverage check itself, rather than the check plus its
// downstream safety net, refuses this range.
func TestRawRangeRejectsAGapAtTheHeadOfTheFightsRange(t *testing.T) {
	d, reportID, objects, _ := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	raw := engine.FixtureLog()
	// Only the back half of the log arrived, but the fight's own
	// recorded start is forced to the very beginning of the file: the
	// one chunk that arrived does not reach back far enough to cover
	// it, which is a gap at the head of the range, not the middle.
	half := int64(len(raw) / 2)
	tail := raw[half:]
	sum := sha256.Sum256(tail)
	if _, err := d.Reports.PutRawChunk(ctx, reportID, reports.RawChunk{
		Start: half, End: int64(len(raw)), SHA256: sum[:],
	}); err != nil {
		t.Fatal(err)
	}
	if err := objects.Put(ctx, store.Keys{ReportID: reportID}.Raw(half), pack(t, tail), store.PutOptions{}); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Reports.Pool.Exec(ctx,
		`update fights set raw_start_offset = 0 where report_id = $1`, reportID); err != nil {
		t.Fatal(err)
	}

	fights, err := d.Reports.Fights(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	chunks, err := d.Reports.RawChunks(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok, err := d.rawRange(ctx, reportID, chunks, fights[0]); err != nil || ok {
		t.Fatalf("rawRange = ok %v, err %v; want ok=false: a chunk starting after the fight's "+
			"own recorded start is a gap at the head of the range, not full coverage", ok, err)
	}
}

func TestSampleFindsNoFightWhenTheRawBytesGarbleTheOpeningLine(t *testing.T) {
	d, reportID, objects, rank := fixtureUpload(t, false)
	ctx := context.Background()
	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	fights, err := d.Reports.Fights(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	start, _, err := d.Reports.FightRawRange(ctx, reportID, fights[0].Index)
	if err != nil {
		t.Fatal(err)
	}
	if start == nil {
		t.Fatal("the fixture fight should have a recorded raw range")
	}
	raw := engine.FixtureLog()
	// The raw chunk claims to cover the whole report from offset zero,
	// but the bytes behind it actually begin a few characters into the
	// fight's own opening line: a stand-in for a companion chunk
	// boundary that lands mid-line. The re-parse should find this
	// inconclusive, not tampered.
	garbled := raw[int(*start)+5:]
	sum := sha256.Sum256(garbled)
	if _, err := d.Reports.PutRawChunk(ctx, reportID, reports.RawChunk{
		Start: 0, End: int64(len(raw)), SHA256: sum[:],
	}); err != nil {
		t.Fatal(err)
	}
	if err := objects.Put(ctx, store.Keys{ReportID: reportID}.Raw(0), pack(t, garbled), store.PutOptions{}); err != nil {
		t.Fatal(err)
	}

	if err := Sample(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged != nil {
		t.Fatalf("a re-parse that found no fight was flagged %q", *rep.Flagged)
	}
	if len(rank.removed) != 0 {
		t.Fatal("nothing should have been withdrawn from a re-parse that found no fight")
	}
}

func TestScheduleNeverPanicsWhenRacingClose(t *testing.T) {
	d, _, _, _ := fixtureUpload(t, false)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w := NewWorker(d)
	done := make(chan struct{})
	go func() { w.Run(ctx); close(done) }()

	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Schedule("nosuchreport")
		}()
	}
	w.Close()
	wg.Wait()
	<-done
}

func TestSampleOnAnUnknownReport(t *testing.T) {
	d, _, _, _ := fixtureUpload(t, false)
	if err := Sample(context.Background(), d, "nosuchreport"); err == nil {
		t.Fatal("an unknown report must report itself")
	}
}

func TestLayoutOfReadsTheReportsHealth(t *testing.T) {
	if got := layoutOf(reports.Report{}); got != "" {
		t.Fatalf("layout = %q, want empty", got)
	}
	if got := layoutOf(reports.Report{Health: []byte("not json")}); got != "" {
		t.Fatalf("layout = %q, want empty for unreadable health", got)
	}
	if got := layoutOf(reports.Report{Health: []byte(`{"layout":"retail-v16"}`)}); got != "retail-v16" {
		t.Fatalf("layout = %q", got)
	}
}

func TestSessionOptionsForPinsAKnownLayout(t *testing.T) {
	o := engine.SessionOptionsFor("r", engine.FixtureBase, true, "retail-v16")
	if o.Layout.Name == "" {
		t.Fatal("a known layout should be pinned")
	}
	o = engine.SessionOptionsFor("r", engine.FixtureBase, true, "no-such-layout")
	if o.Layout.Name != "" {
		t.Fatalf("layout = %q, want the session left to infer", o.Layout.Name)
	}
}

// Run used to return on ctx.Done, and ctx is the signal context: on
// SIGTERM the consumer exited before http.Server.Shutdown drained the
// handlers, so a completion call arriving during shutdown queued a
// check nobody would ever read. The consumer has to outlive the
// handlers, and the check itself has to survive the signal that
// started the shutdown.
func TestAReportScheduledDuringShutdownIsStillChecked(t *testing.T) {
	d, reportID, objects, rank := fixtureUpload(t, false)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := Report(context.Background(), d, reportID); err != nil {
		t.Fatal(err)
	}
	seedRawChunk(t, d, objects, reportID, engine.FixtureLog())
	trimStoredEvents(t, d, objects, reportID)

	w := NewWorker(d)
	done := make(chan struct{})
	go func() { w.Run(ctx); close(done) }()

	// The signal lands first, exactly as it does in cmd/api: the
	// listener stops, the handlers keep draining, and this is one of
	// their completion calls.
	cancel()
	w.Schedule(reportID)
	w.Close()
	<-done

	rep, err := d.Reports.Get(context.Background(), reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged == nil || *rep.Flagged != FlagTampered {
		t.Fatalf("flagged = %v, want the check to have run despite the signal", rep.Flagged)
	}
	if len(rank.removed) != 1 || rank.removed[0] != reportID {
		t.Fatalf("removed = %v, want the report withdrawn", rank.removed)
	}
}

// trimStoredEvents rewrites a report's first fight with one event
// missing, which is what a companion that trimmed its own deaths would
// have uploaded.
func trimStoredEvents(t *testing.T, d Deps, objects *memObjects, reportID string) {
	t.Helper()
	ctx := context.Background()
	fights, err := d.Reports.Fights(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	key := store.Keys{ReportID: reportID}.FightEvents(fights[0].Index)
	stored, _ := objects.get(key)
	events, err := logparquet.Unmarshal(stored)
	if err != nil {
		t.Fatal(err)
	}
	trimmed, err := logparquet.Marshal(events[:len(events)-1])
	if err != nil {
		t.Fatal(err)
	}
	if err := objects.Put(ctx, key, trimmed, store.PutOptions{}); err != nil {
		t.Fatal(err)
	}
}
