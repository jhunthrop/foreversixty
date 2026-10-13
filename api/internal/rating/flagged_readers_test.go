package rating

// Task L6.12: a report flagged tampered or unverified counts nowhere a player looks, and the
// backfill stops rating one. The rating endpoints' character history (ReadCharacterRating:
// the fight cards and trend) and the backfill's pick of unrated fights (staleFights) select
// their reports through presence.EvidenceReport.

import (
	"context"
	"testing"
	"time"
)

// The character's rating history never lists a flagged report's card.
func TestReadCharacterRatingLeavesOutAFlaggedReport(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	flagged := sv2Report + "c"
	inReport := func() int {
		t.Helper()
		rows, _, err := fx.store.ReadCharacterRating(ctx, fx.main, characterTrendPageSize, nil)
		if err != nil {
			t.Fatal(err)
		}
		n := 0
		for _, r := range rows {
			if r.ReportID == flagged {
				n++
			}
		}
		return n
	}
	if n := inReport(); n != 4 {
		t.Fatalf("before the flag the history lists %d of the report's cards, want its 4", n)
	}
	flagReport(t, fx, flagged)
	if n := inReport(); n != 0 {
		t.Errorf("the history lists %d cards of a flagged report, want none", n)
	}
}

// The backfill never rates a flagged report's unrated fights; the same fight in an
// unflagged report is picked.
func TestStaleFightsSkipsAFlaggedReportsUnratedFights(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	t.Cleanup(func() {
		pool.Exec(context.Background(), `delete from reports where id like '000l612-%'`)
	})
	at := time.Now().Add(-24 * time.Hour)
	for _, r := range []struct{ id, flagged string }{{"000l612-clean", ""}, {"000l612-flagged", "tampered"}} {
		mustExec(t, pool, `insert into reports (id, title, visibility, zone, status, flagged, logging_character)
			values ($1, 'L6.12', 'public', 'Molten Core', 'complete', nullif($2, ''), 'us/normal/l612logger')`,
			r.id, r.flagged)
		mustExec(t, pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
			values ($1, 1, 990699, 'L6.12 boss', 3, true, $2)`, r.id, at.UnixMilli())
	}
	picked, err := (&Store{Pool: pool}).staleFights(ctx, BackfillBatchSize)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, sf := range picked {
		got[sf.ReportID] = true
	}
	if !got["000l612-clean"] {
		t.Fatalf("the clean report's unrated fight was not picked: %+v", picked)
	}
	if got["000l612-flagged"] {
		t.Errorf("the flagged report's unrated fight was picked for rating")
	}
}
