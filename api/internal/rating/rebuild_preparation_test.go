// api/internal/rating/rebuild_preparation_test.go
package rating

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/digest"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	rebuildEncounter  = 990003
	rebuildFlaskSpell = 17626
	rebuildOldModel   = "rating-2020-01-01"
)

// rebuildFight is a fight whose one dps carries a seeded flask, on a private encounter so
// no other test's digest shares its bracket.
func rebuildFight(id string, kill bool) RatedFight {
	f := fightFixture(id, kill, summary.RosterRow{GUID: "g1", Name: "Flasked", Class: "Warrior", Spec: "Fury", Role: "dps"})
	f.EncounterID = rebuildEncounter
	f.Summary.EncounterID = rebuildEncounter
	f.Summary.Combatants = []summary.CombatantRow{{GUID: "g1", Gear: []event.Item{{ID: 1}}}}
	f.Summary.Auras = []summary.AuraTrack{{TargetGUID: "g1", SpellID: rebuildFlaskSpell, Seeded: true,
		Segments: []summary.Segment{{StartMS: 0, EndMS: f.Summary.DurationMS}}}}
	return f
}

func rebuildWant(t *testing.T, f RatedFight) float64 {
	t.Helper()
	rc, _ := consumables.Load().For("dps")
	v, ok := ratingengine.PreparationValue(f.Summary, "g1", &rc)
	if !ok || v <= 0 {
		t.Fatalf("fixture preparation value = %v, %v; want a positive value", v, ok)
	}
	return v
}

// seedStaleRebuild rates the fights, ages their rows to the old model and plants a wrong
// preparation digest holding two zeros; it returns the getter serving the summaries.
func seedStaleRebuild(t *testing.T, pool *pgxpool.Pool, fights ...RatedFight) fakeGetter {
	t.Helper()
	ctx := context.Background()
	pool.Exec(ctx, `delete from rating_scores`)
	pool.Exec(ctx, `delete from rating_percentile_digests where component = 'preparation'`)
	objects := map[string][]byte{}
	for _, f := range fights {
		if err := (&Store{Pool: pool}).RateFight(ctx, f); err != nil {
			t.Fatal(err)
		}
		body, err := json.Marshal(f.Summary)
		if err != nil {
			t.Fatal(err)
		}
		objects[store.Keys{ReportID: f.ReportID}.FightSummary(f.FightIndex)] = body
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where encounter_id = $1`, rebuildEncounter) })
	pool.Exec(ctx, `update rating_scores set model_version = $1 where encounter_id = $2`, rebuildOldModel, rebuildEncounter)
	wrong := digest.New()
	wrong.Add(0)
	wrong.Add(0)
	raw, _ := wrong.MarshalBinary()
	if _, err := pool.Exec(ctx, `update rating_percentile_digests set digest = $1, n = 2 where component = 'preparation'`, raw); err != nil {
		t.Fatal(err)
	}
	return fakeGetter{objects: objects}
}

func preparationDigest(t *testing.T, pool *pgxpool.Pool) (int64, *digest.Digest) {
	t.Helper()
	var raw []byte
	var n int64
	err := pool.QueryRow(context.Background(),
		`select digest, n from rating_percentile_digests where encounter_id = $1 and component = 'preparation' and kill`,
		rebuildEncounter).Scan(&raw, &n)
	if err != nil {
		t.Fatal(err)
	}
	d, err := digest.Unmarshal(raw)
	if err != nil {
		t.Fatal(err)
	}
	return n, d
}

func TestRebuildPreparation_RefoldsFromSeededAuras(t *testing.T) {
	pool := testPool(t)
	a, b := rebuildFight("rebuild-a", true), rebuildFight("rebuild-b", true)
	getter := seedStaleRebuild(t, pool, a, b)
	counts, err := RebuildPreparation(context.Background(), BackfillDeps{Store: &Store{Pool: pool}, Summaries: getter})
	if err != nil {
		t.Fatal(err)
	}
	if want := (RebuildCounts{Folded: 2, Rerated: 2, RolledUp: 2}); counts != want {
		t.Fatalf("counts = %+v, want %+v", counts, want)
	}
	n, d := preparationDigest(t, pool)
	if want := rebuildWant(t, a); n != 2 || d.Count() != 2 || d.Quantile(0.5) != want {
		t.Errorf("digest n=%d count=%d median=%v, want 2 values of %v", n, d.Count(), d.Quantile(0.5), want)
	}
	var stale int
	pool.QueryRow(context.Background(), `select count(*) from rating_scores where encounter_id = $1 and model_version <> $2`,
		rebuildEncounter, ratingengine.DefaultModelVersion).Scan(&stale)
	if stale != 0 {
		t.Errorf("%d rows still off the current model version", stale)
	}
}

func TestRebuildPreparation_IsIdempotent(t *testing.T) {
	pool := testPool(t)
	getter := seedStaleRebuild(t, pool, rebuildFight("rebuild-a", true), rebuildFight("rebuild-b", true))
	deps := BackfillDeps{Store: &Store{Pool: pool}, Summaries: getter}
	if _, err := RebuildPreparation(context.Background(), deps); err != nil {
		t.Fatal(err)
	}
	second, err := RebuildPreparation(context.Background(), deps)
	if err != nil {
		t.Fatal(err)
	}
	if second.Rerated != 0 || second.Folded != 2 {
		t.Errorf("second run = %+v, want Rerated 0 and Folded 2 from an empty digest", second)
	}
	if n, _ := preparationDigest(t, pool); n != 2 {
		t.Errorf("digest n = %d after two runs, want 2", n)
	}
}

func TestRebuildPreparation_WipesAreNotFolded(t *testing.T) {
	pool := testPool(t)
	getter := seedStaleRebuild(t, pool, rebuildFight("rebuild-a", true), rebuildFight("rebuild-wipe", false))
	if _, err := RebuildPreparation(context.Background(), BackfillDeps{Store: &Store{Pool: pool}, Summaries: getter}); err != nil {
		t.Fatal(err)
	}
	if n, _ := preparationDigest(t, pool); n != 1 {
		t.Errorf("digest n = %d, want 1: the wipe's preparation must not enter it", n)
	}
}

func TestRebuildPreparation_DryRunWritesNothing(t *testing.T) {
	pool := testPool(t)
	getter := seedStaleRebuild(t, pool, rebuildFight("rebuild-a", true), rebuildFight("rebuild-b", true))
	counts, err := RebuildPreparationWith(context.Background(),
		BackfillDeps{Store: &Store{Pool: pool}, Summaries: getter}, RebuildOptions{DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if counts.Folded != 2 || counts.Rerated != 2 || counts.RolledUp != 0 {
		t.Errorf("dry-run counts = %+v, want Folded 2, Rerated 2 (would), RolledUp 0", counts)
	}
	n, d := preparationDigest(t, pool)
	if n != 2 || d.Quantile(0.5) != 0 {
		t.Errorf("dry run changed the digest: n=%d median=%v", n, d.Quantile(0.5))
	}
	var stale int
	pool.QueryRow(context.Background(), `select count(*) from rating_scores where model_version = $1`, rebuildOldModel).Scan(&stale)
	if stale != 2 {
		t.Errorf("dry run re-rated rows: %d still old, want 2", stale)
	}
}

func TestRebuildPreparation_ResumesPastTheCursorWithoutClearing(t *testing.T) {
	pool := testPool(t)
	getter := seedStaleRebuild(t, pool, rebuildFight("rebuild-a", true), rebuildFight("rebuild-b", true))
	counts, err := RebuildPreparationWith(context.Background(),
		BackfillDeps{Store: &Store{Pool: pool}, Summaries: getter},
		RebuildOptions{After: &FightCursor{ReportID: "rebuild-a", FightIndex: 1}})
	if err != nil {
		t.Fatal(err)
	}
	if counts.Folded != 1 {
		t.Errorf("Folded = %d past the cursor, want 1", counts.Folded)
	}
	if n, _ := preparationDigest(t, pool); n != 3 {
		t.Errorf("digest n = %d, want 3: the planted 2 kept (no clear on resume) plus 1 folded", n)
	}
}

func TestRebuildPreparation_DryRunInsertsNoEmptyDigestRows(t *testing.T) {
	pool := testPool(t)
	getter := seedStaleRebuild(t, pool, rebuildFight("rebuild-a", true))
	if _, err := pool.Exec(context.Background(),
		`delete from rating_percentile_digests where component = 'preparation'`); err != nil {
		t.Fatal(err)
	}
	if _, err := RebuildPreparationWith(context.Background(),
		BackfillDeps{Store: &Store{Pool: pool}, Summaries: getter}, RebuildOptions{DryRun: true}); err != nil {
		t.Fatal(err)
	}
	var rows int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from rating_percentile_digests where component = 'preparation'`).Scan(&rows); err != nil {
		t.Fatal(err)
	}
	if rows != 0 {
		t.Errorf("a dry run left %d preparation digest rows, want 0 (read-only)", rows)
	}
}
