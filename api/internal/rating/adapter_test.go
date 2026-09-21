// api/internal/rating/adapter_test.go
package rating

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := testURL(t) // see Step 2: this task also adds testURL to this package
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func TestPlacementReadsEmptyBracketAsNotOK(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	a := &percentileAdapter{Tx: tx, Fold: true}
	b := ratingengine.Bracket{EncounterID: 111111, Difficulty: 1, Spec: "warrior-fury", Role: "dps",
		KillTimeBand: "typical", Kill: true, Component: ratingengine.ComponentNameOutput}
	_, _, ok := a.Placement(b, 50)
	if ok {
		t.Error("a never-seen bracket must report ok = false")
	}
}

func TestPlacementFoldsOnKillForAKillGatedComponent(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 222222, Difficulty: 1, Spec: "mage-frost", Role: "dps",
		KillTimeBand: "typical", Kill: true, Component: ratingengine.ComponentNameOutput}
	a := &percentileAdapter{Tx: tx, Fold: true}
	for i, v := range []float64{10, 20, 30, 40, 50} {
		pct, n, ok := a.Placement(b, v)
		if i == 0 && ok {
			t.Error("the first fold into a brand-new bracket must still report ok = false (nothing to compare against yet)")
		}
		if i > 0 && (!ok || n != int64(i)) {
			t.Errorf("fold %d: pct=%v n=%v ok=%v, want ok and n=%d", i, pct, n, ok, i)
		}
	}
}

func TestPlacementDoesNotFoldAWipeIntoAKillGatedComponent(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 333333, Difficulty: 1, Spec: "priest-holy", Role: "healer",
		KillTimeBand: "", Kill: false, Component: ratingengine.ComponentNameActivity}
	a := &percentileAdapter{Tx: tx, Fold: true}
	if _, _, ok := a.Placement(b, 88); ok {
		t.Fatal("first call on an empty bracket must be ok=false")
	}
	// A second call with the same wipe bracket must see the SAME empty state: nothing
	// was folded, because Activity is not Survival/Mechanics and this bracket is a wipe.
	if _, n, ok := a.Placement(b, 88); ok || n != 0 {
		t.Errorf("wipe fold into a kill-gated component: ok=%v n=%v, want ok=false n=0", ok, n)
	}
}

func TestPlacementFoldsAWipeIntoASurvivalSubBracket(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 444444, Difficulty: 1, Spec: "warrior-protection", Role: "tank",
		KillTimeBand: "", Kill: false, Component: ratingengine.ComponentSurvivalAvoidableHit}
	a := &percentileAdapter{Tx: tx, Fold: true}
	a.Placement(b, 10)
	_, n, ok := a.Placement(b, 20)
	if !ok || n != 1 {
		t.Errorf("Survival must fold on a wipe: ok=%v n=%v, want ok=true n=1", ok, n)
	}
}

func TestPlacementNeverFoldsWhenFoldIsFalse(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 555555, Difficulty: 1, Spec: "hunter-survival", Role: "dps",
		KillTimeBand: "typical", Kill: true, Component: ratingengine.ComponentNameOutput}
	fold := &percentileAdapter{Tx: tx, Fold: true}
	fold.Placement(b, 100)
	readOnly := &percentileAdapter{Tx: tx, Fold: false}
	readOnly.Placement(b, 999)
	_, n, _ := fold.Placement(b, 1) // a fresh read confirms the read-only call added nothing
	if n != 1 {
		t.Errorf("n = %d after a read-only call, want 1 (unchanged)", n)
	}
}

func TestKillTimeBandNeverFolds(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	a := &percentileAdapter{Tx: tx, Fold: true}
	if _, ok := a.KillTimeBand(666666, 1, 180000); ok {
		t.Fatal("a never-seen encounter/difficulty must report ok = false")
	}
	if err := foldKillDuration(ctx, tx, 666666, 1, 180000); err != nil {
		t.Fatal(err)
	}
	band, ok := a.KillTimeBand(666666, 1, 180000)
	if !ok || band != ratingengine.BandTypical {
		t.Errorf("band = %q ok=%v, want typical/true after exactly one fold", band, ok)
	}
	// KillTimeBand itself must not have folded a second time.
	if err := foldKillDuration(ctx, tx, 666666, 1, 90000); err != nil { // a much faster kill
		t.Fatal(err)
	}
	band, _ = a.KillTimeBand(666666, 1, 90000)
	if band != ratingengine.BandFast {
		t.Errorf("band = %q, want fast (median should now reflect both folds, not be stuck on a single call's snapshot)", band)
	}
}
