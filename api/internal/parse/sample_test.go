package parse

import (
	"context"
	"crypto/sha256"
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

	// Rewrite the stored events with one event missing, which is what a
	// companion that trimmed its own deaths would have uploaded.
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
