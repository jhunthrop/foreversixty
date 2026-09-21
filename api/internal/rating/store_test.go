// api/internal/rating/store_test.go
package rating

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func fightFixture(reportID string, kill bool, players ...summary.RosterRow) RatedFight {
	return RatedFight{
		ReportID: reportID, FightIndex: 1, Region: "us", Ruleset: "normal",
		FoughtAt: time.Date(2026, 12, 9, 1, 0, 0, 0, time.UTC), EncounterID: 667,
		Summary: summary.Summary{
			EncounterID: 667, Difficulty: 3, DurationMS: 180000, Kill: kill,
			Roster: players,
		},
	}
}

func TestRateFightWritesOneRowPerRosterPlayer(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-test-1", true,
		summary.RosterRow{GUID: "g1", Name: "Simfury", Class: "Warrior", Spec: "Fury", Role: "dps"},
		summary.RosterRow{GUID: "g2", Name: "Healface", Class: "Priest", Spec: "Holy", Role: "healer"},
	)
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })

	var n int
	if err := pool.QueryRow(ctx, `select count(*) from rating_scores where report_id = $1`, f.ReportID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("wrote %d rows, want 2", n)
	}
	var key string
	if err := pool.QueryRow(ctx, `select player_key from rating_scores where report_id = $1 and player_name = 'Simfury'`, f.ReportID).Scan(&key); err != nil {
		t.Fatal(err)
	}
	if key != "us/normal/simfury" {
		t.Errorf("player_key = %q, want us/normal/simfury", key)
	}
}

func TestRateFightIsIdempotentAndDoesNotDoubleFoldDigests(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-test-idempotent", true,
		summary.RosterRow{GUID: "g1", Name: "Repeatme", Class: "Warrior", Spec: "Fury", Role: "dps"},
	)
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID)
	})
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	var nAfterFirst int64
	pool.QueryRow(ctx,
		`select n from rating_percentile_digests where encounter_id = 667 and difficulty = 3
		 and spec = 'warrior-fury' and role = 'dps' and component = 'output'`).Scan(&nAfterFirst)

	if err := store.RateFight(ctx, f); err != nil { // a re-ingest of the same fight
		t.Fatal(err)
	}
	var rowCount int
	pool.QueryRow(ctx, `select count(*) from rating_scores where report_id = $1`, f.ReportID).Scan(&rowCount)
	if rowCount != 1 {
		t.Fatalf("rewriting the same fight produced %d rows, want exactly 1 (upsert, not a duplicate)", rowCount)
	}
	var nAfterSecond int64
	pool.QueryRow(ctx,
		`select n from rating_percentile_digests where encounter_id = 667 and difficulty = 3
		 and spec = 'warrior-fury' and role = 'dps' and component = 'output'`).Scan(&nAfterSecond)
	if nAfterSecond != nAfterFirst {
		t.Errorf("digest n grew from %d to %d on a rewrite - the same value was folded twice", nAfterFirst, nAfterSecond)
	}
}

func TestRateFightStoresSixComponentsPerPlayer(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-test-components", true,
		summary.RosterRow{GUID: "g1", Name: "Sixparts", Class: "Mage", Spec: "Frost", Role: "dps"},
	)
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	var raw json.RawMessage
	if err := pool.QueryRow(ctx, `select components from rating_scores where report_id = $1`, f.ReportID).Scan(&raw); err != nil {
		t.Fatal(err)
	}
	var dtos []componentDTO
	if err := json.Unmarshal(raw, &dtos); err != nil {
		t.Fatal(err)
	}
	if len(dtos) != 6 {
		t.Fatalf("stored %d components, want 6", len(dtos))
	}
}
