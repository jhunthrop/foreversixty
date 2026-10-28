package rating

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

type recordingHook struct {
	fight RatedFight
	cards []CardRow
	err   error
}

func (h *recordingHook) InTx(_ context.Context, _ pgx.Tx, f RatedFight, cards []CardRow) error {
	h.fight, h.cards = f, cards
	return h.err
}

func hookFight(reportID string) RatedFight {
	return fightFixture(reportID, true,
		summary.RosterRow{GUID: "g1", Name: "Hookfury", Class: "Warrior", Spec: "Fury", Role: "dps"},
		summary.RosterRow{GUID: "g2", Name: "Hookheal", Class: "Priest", Spec: "Holy", Role: "healer"},
	)
}

func countScores(t *testing.T, store *Store, reportID string) int {
	t.Helper()
	var n int
	if err := store.Pool.QueryRow(context.Background(),
		`select count(*) from rating_scores where report_id = $1`, reportID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestRateFightCallsTheHookWithEveryCardRow(t *testing.T) {
	pool := testPool(t)
	hook := &recordingHook{}
	store := &Store{Pool: pool, Hook: hook}
	f := hookFight("hook-test-1")
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	if hook.fight.ReportID != f.ReportID || len(hook.cards) != 2 {
		t.Fatalf("hook saw %q with %d cards, want %q with 2", hook.fight.ReportID, len(hook.cards), f.ReportID)
	}
	for _, c := range hook.cards {
		if c.PlayerKey == "" || c.PlayerName == "" {
			t.Errorf("card %+v lacks its player key or name", c)
		}
	}
}

func TestRateFightHookErrorRollsTheRatingBack(t *testing.T) {
	pool := testPool(t)
	boom := errors.New("hook failed")
	store := &Store{Pool: pool, Hook: &recordingHook{err: boom}}
	f := hookFight("hook-test-2")
	err := store.RateFight(context.Background(), f)
	if !errors.Is(err, boom) {
		t.Fatalf("RateFight = %v, want the hook's error", err)
	}
	if n := countScores(t, store, f.ReportID); n != 0 {
		t.Fatalf("%d rating rows survived a failed hook", n)
	}
}

func TestRateFightWithoutAHookWritesAsBefore(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	f := hookFight("hook-test-3")
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	if n := countScores(t, store, f.ReportID); n != 2 {
		t.Fatalf("wrote %d rows, want 2", n)
	}
}
