package events

import (
	"context"
	"testing"
)

// Every row has an icon: its producer's, else its group's; a row stored without one (written
// before icons) reads its group's too.
func TestActivityRowsAlwaysCarryAnIcon(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "icons")
	ctx := context.Background()
	for _, e := range []Event{
		{Kind: GuildSimPassed, DedupeKey: "a", Title: "A", Icon: "ability_warrior_innerrage"},
		{Kind: RaidShort, DedupeKey: "b", Title: "B"},
	} {
		if err := st.Record(ctx, []int64{uid}, e); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx, `insert into user_activity (user_id, kind, title, body, link, dedupe_key)
		values ($1, 'sim.finished', 'Old', '', '', 'old')`, uid); err != nil {
		t.Fatal(err)
	}
	rows, _, err := st.Activity(ctx, uid, ActivityQuery{Limit: 10})
	if err != nil || len(rows) != 3 {
		t.Fatalf("rows %+v %v", rows, err)
	}
	want := []string{IconSim, "inv_scroll_08", "ability_warrior_innerrage"}
	for i, r := range rows {
		if r.Icon != want[i] {
			t.Errorf("row %d (%s) icon %q, want %q", i, r.Kind, r.Icon, want[i])
		}
	}
}

func TestEveryGroupHasAnIcon(t *testing.T) {
	for _, g := range Groups {
		if groupIcons[g] == "" {
			t.Errorf("%s has no icon", g)
		}
	}
	if IconOf("no.such.kind", "") != groupIcons[GroupAccount] {
		t.Fatal("an unknown kind reads as account")
	}
}
