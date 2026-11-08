package events

import (
	"context"
	"testing"
)

// A row keeps its producer's icon; a row recorded with none reads as null (the apps' group icon).
func TestActivityRowsCarryTheirIcon(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "icons")
	ctx := context.Background()
	for _, e := range []Event{
		{Kind: GuildSimPassed, DedupeKey: "a", Title: "A", Icon: CrestIcon("Warrior")},
		{Kind: SimFinished, DedupeKey: "b", Title: "B"},
	} {
		if err := st.Record(ctx, []int64{uid}, e); err != nil {
			t.Fatal(err)
		}
	}
	rows, _, err := st.Activity(ctx, uid, ActivityQuery{Limit: 10})
	if err != nil || len(rows) != 2 {
		t.Fatalf("rows %+v %v", rows, err)
	}
	if rows[0].Icon != nil || rows[1].Icon == nil || *rows[1].Icon != "crest:warrior" {
		t.Fatalf("icons %v %v", rows[0].Icon, rows[1].Icon)
	}
}
