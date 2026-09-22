// api/internal/rating/backfill_test.go
package rating

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"testing"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

var errFakeGetterMiss = errors.New("fake getter: no such object")

type fakeGetter struct{ objects map[string][]byte }

func (g fakeGetter) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	b, ok := g.objects[key]
	if !ok {
		return nil, errFakeGetterMiss
	}
	return io.NopCloser(bytes.NewReader(b)), nil
}

func TestBackfillRecomputesStaleRowsAndAdvancesModelVersion(t *testing.T) {
	pool := testPool(t)
	// Named ratingStore, not store: this file also imports logs/engine/store for
	// store.Keys below, and a local variable named store would shadow that package
	// identifier for the rest of the function.
	ratingStore := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("backfill-test-1", true,
		summary.RosterRow{GUID: "g1", Name: "Stale", Class: "Shaman", Spec: "Enhancement", Role: "dps"},
	)
	if err := ratingStore.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })
	// Force the row stale, as if an older model_version had written it.
	if _, err := pool.Exec(ctx, `update rating_scores set model_version = 'rating-2020-01-01' where report_id = $1`, f.ReportID); err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(f.Summary)
	if err != nil {
		t.Fatal(err)
	}
	getter := fakeGetter{objects: map[string][]byte{
		store.Keys{ReportID: f.ReportID}.FightSummary(f.FightIndex): body,
	}}
	n, err := Backfill(ctx, BackfillDeps{Store: ratingStore, Summaries: getter}, 10)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("recomputed %d fights, want 1", n)
	}
	var version string
	pool.QueryRow(ctx, `select model_version from rating_scores where report_id = $1`, f.ReportID).Scan(&version)
	if version != ratingengine.DefaultModelVersion {
		t.Errorf("model_version = %q, want %q", version, ratingengine.DefaultModelVersion)
	}
}

// TestBackfillIsBoundedPerRun proves the batch bound with reads that actually succeed:
// with 3 stale fights and a getter that can satisfy all 3, a batchSize-2 run must
// recompute exactly 2 and leave exactly 1 still stale. A getter where every read fails
// (as an earlier version of this test used) cannot tell "the SQL limit was really 2" apart
// from "the limit was ignored and every fight failed anyway" - both report 0 recomputed
// and 3 remaining regardless of the bound, so that shape does not actually test anything
// about boundedness. This version does.
func TestBackfillIsBoundedPerRun(t *testing.T) {
	pool := testPool(t)
	ratingStore := &Store{Pool: pool} // see the naming note in the test above
	ctx := context.Background()
	objects := map[string][]byte{}
	for i := 0; i < 3; i++ {
		f := fightFixture(fightIDFor(i), true,
			summary.RosterRow{GUID: "g1", Name: "Bounded", Class: "Shaman", Spec: "Elemental", Role: "dps"})
		if err := ratingStore.RateFight(ctx, f); err != nil {
			t.Fatal(err)
		}
		pool.Exec(ctx, `update rating_scores set model_version = 'rating-2020-01-01' where report_id = $1`, f.ReportID)
		body, err := json.Marshal(f.Summary)
		if err != nil {
			t.Fatal(err)
		}
		objects[store.Keys{ReportID: f.ReportID}.FightSummary(f.FightIndex)] = body
	}
	t.Cleanup(func() {
		for i := 0; i < 3; i++ {
			pool.Exec(ctx, `delete from rating_scores where report_id = $1`, fightIDFor(i))
		}
	})
	getter := fakeGetter{objects: objects} // every one of the 3 stale fights' summaries is available
	n, err := Backfill(ctx, BackfillDeps{Store: ratingStore, Summaries: getter}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("recomputed %d fights, want exactly 2 (batchSize 2 of 3 stale, all reads succeed)", n)
	}
	remaining, err := ratingStore.staleFights(ctx, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(remaining) != 1 {
		t.Fatalf("staleFights after a bounded run = %d, want 1 (3 stale - 2 recomputed)", len(remaining))
	}
}

func fightIDFor(i int) string {
	return "backfill-bounded-" + string(rune('a'+i))
}
