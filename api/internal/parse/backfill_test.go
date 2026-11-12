package parse

import (
	"context"
	"crypto/sha256"
	"io"
	"log/slog"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// mustNotNotify fails the test the moment anything is announced.
type mustNotNotify struct{ t *testing.T }

func (m mustNotNotify) UploadParsed(_ context.Context, e notify.UploadParsed) {
	m.t.Errorf("the backfill announced report %s; a backfill is silent", e.ReportID)
}

// seedLiveReport seeds a completed live report with no fights, the shape the week of
// 10 Nov 2026 left behind, whose raw chunks hold the fixture log split at cut.
func seedLiveReport(t *testing.T, cut int) (Deps, string, *memObjects, *fakeCompleted) {
	t.Helper()
	pool := testPool(t)
	ctx := context.Background()
	owner, err := (&auth.Store{Pool: pool}).UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	rs := &reports.Store{Pool: pool}
	reportID := auth.NewReportID()
	character := "us/hardcore/baelgrim"
	if _, err := rs.Create(ctx, reports.Report{
		ID: reportID, OwnerID: &owner.ID, Title: "Tuesday", Visibility: reports.Public,
		Status: reports.StatusComplete, LoggingCharacter: &character,
	}); err != nil {
		t.Fatal(err)
	}
	objects := newMemObjects()
	raw := engine.FixtureLog()
	for _, part := range [][2]int{{0, cut}, {cut, len(raw)}} {
		body := raw[part[0]:part[1]]
		sum := sha256.Sum256(body)
		if _, err := rs.PutRawChunk(ctx, reportID, reports.RawChunk{
			Start: int64(part[0]), End: int64(part[1]), SHA256: sum[:],
		}); err != nil {
			t.Fatal(err)
		}
		if err := objects.Put(ctx, store.Keys{ReportID: reportID}.Raw(int64(part[0])), pack(t, body), store.PutOptions{}); err != nil {
			t.Fatal(err)
		}
	}
	completed := &fakeCompleted{}
	return Deps{
		Reports: rs, Objects: objects, Rank: &fakeRanker{}, Completed: completed,
		Notify: mustNotNotify{t}, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}, reportID, objects, completed
}

// The dry run counts the fights the raw chunks hold and writes nothing; the applied run
// writes them through the live path's own writeFight, completes the report, tells the
// completion hook, and announces nothing.
func TestBackfillCountsThenWritesTheFightsARawStreamHolds(t *testing.T) {
	d, reportID, objects, completed := seedLiveReport(t, len(engine.FixtureLog())/2)
	ctx := context.Background()

	dry, err := Backfill(ctx, d, reportID, false)
	if err != nil {
		t.Fatal(err)
	}
	if dry.Fights != 1 || dry.Encounters != 1 || dry.Kills != 1 || dry.Chunks != 2 || dry.GapAt != nil {
		t.Fatalf("dry run = %+v, want one killed encounter from two gapless chunks", dry)
	}
	if dry.Health.Lines != int64(engine.FixtureLines()) || dry.Health.MissingHeader {
		t.Fatalf("dry-run health = %+v", dry.Health)
	}
	if fights, err := d.Reports.Fights(ctx, reportID); err != nil || len(fights) != 0 {
		t.Fatalf("the dry run wrote %d fights (%v)", len(fights), err)
	}
	if _, ok := objects.get(store.Keys{ReportID: reportID}.Report()); ok {
		t.Fatal("the dry run wrote report.json")
	}

	applied, err := Backfill(ctx, d, reportID, true)
	if err != nil {
		t.Fatal(err)
	}
	if applied.Fights != dry.Fights {
		t.Fatalf("applied %d fights, the dry run counted %d", applied.Fights, dry.Fights)
	}
	fights, err := d.Reports.Fights(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if len(fights) != 1 || !fights[0].Verified || fights[0].Name != "Warden Kelthas" {
		t.Fatalf("fights = %+v", fights)
	}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Status != reports.StatusComplete || rep.EngineVersion != engine.Version || len(rep.Health) == 0 {
		t.Fatalf("report = %+v", rep)
	}
	keys := store.Keys{ReportID: reportID}
	for _, key := range []string{keys.Report(), keys.FightSummary(1), keys.FightEvents(1)} {
		if _, ok := objects.get(key); !ok {
			t.Errorf("object %s was not written", key)
		}
	}
	if got := completed.all(); len(got) != 1 || got[0] != reportID {
		t.Fatalf("completion hook told of %v, want %s once", got, reportID)
	}
	if len(d.Rank.(*fakeRanker).written) != 1 {
		t.Fatal("the backfilled encounter was not ranked")
	}

	// A second run finds the fights and refuses rather than overwriting them.
	if _, err := Backfill(ctx, d, reportID, true); err == nil {
		t.Fatal("a report that holds fights must be refused")
	}
}

// A hole in the raw stream ends the parse there and is reported, rather than feeding the
// engine bytes that do not follow one another.
func TestBackfillStopsAtAGapAndSaysWhere(t *testing.T) {
	d, reportID, _, _ := seedLiveReport(t, len(engine.FixtureLog())/2)
	ctx := context.Background()
	half := int64(len(engine.FixtureLog()) / 2)
	if err := d.Reports.DeleteRawChunk(ctx, reportID, 0); err != nil {
		t.Fatal(err)
	}
	res, err := Backfill(ctx, d, reportID, false)
	if err != nil {
		t.Fatal(err)
	}
	if res.GapAt == nil || *res.GapAt != 0 || res.Chunks != 0 || res.Fights != 0 {
		t.Fatalf("result = %+v, want a gap at 0 and nothing streamed (second chunk starts at %d)", res, half)
	}
}

// A report with no raw chunks has nothing to backfill from and says so.
func TestBackfillRefusesAReportWithoutRawChunks(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	if _, err := Backfill(context.Background(), d, reportID, false); err == nil {
		t.Fatal("a report without raw chunks must be refused")
	}
}
