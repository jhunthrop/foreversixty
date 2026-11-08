package events

import (
	"context"
	"fmt"
	"slices"
	"testing"
	"time"
)

func activityIDs(rows []ActivityRow) []int64 {
	out := make([]int64, len(rows))
	for i, r := range rows {
		out[i] = r.ID
	}
	return out
}

func activityKinds(rows []ActivityRow) []Kind {
	out := make([]Kind, len(rows))
	for i, r := range rows {
		out[i] = r.Kind
	}
	return out
}

// recordAll records one event of each kind for uid, in order, with distinct keys.
func recordAll(t *testing.T, st *Store, uid int64, kinds ...Kind) {
	t.Helper()
	for i, k := range kinds {
		e := Event{Kind: k, DedupeKey: fmt.Sprintf("%s:%d", k, i), Title: "T " + string(k), Body: "B", Link: "https://foreversixty.gg/x"}
		if err := st.Record(context.Background(), []int64{uid}, e); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRecordKeepsOneRowPerRecipientAndKey(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	a, b := seedUser(t, pool, "a"), seedUser(t, pool, "b")
	ctx := context.Background()
	e := Event{Kind: RaidPosted, DedupeKey: "raid.posted:1", Title: "MC posted", Body: "Thursday", Link: "https://foreversixty.gg/r/1"}
	for range 2 {
		if err := st.Record(ctx, []int64{a, b}, e); err != nil {
			t.Fatal(err)
		}
	}
	if err := st.Record(ctx, nil, e); err != nil {
		t.Fatal(err)
	}
	for _, uid := range []int64{a, b} {
		rows, next, err := st.Activity(ctx, uid, ActivityQuery{Limit: 30})
		if err != nil || len(rows) != 1 || next != nil {
			t.Fatalf("rows %+v next %v err %v", rows, next, err)
		}
		r := rows[0]
		if r.Kind != RaidPosted || r.Group != GroupRaids || r.Title != e.Title || r.Body != e.Body || r.Link != e.Link || r.Read || r.OccurredAt.IsZero() {
			t.Fatalf("row %+v", r)
		}
	}
}

func TestActivityPagesNewestFirstWithACursor(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid, other := seedUser(t, pool, "p"), seedUser(t, pool, "o")
	ctx := context.Background()
	recordAll(t, st, uid, RaidPosted, SimFinished, ReportReady, RaidShort, CharacterLinked)
	recordAll(t, st, other, SimFinished)
	first, next, err := st.Activity(ctx, uid, ActivityQuery{Limit: 2})
	if err != nil || next == nil || !slices.Equal(activityKinds(first), []Kind{CharacterLinked, RaidShort}) {
		t.Fatalf("first %v next %v err %v", activityKinds(first), next, err)
	}
	if *next != first[1].ID {
		t.Fatalf("next %d", *next)
	}
	second, next2, _ := st.Activity(ctx, uid, ActivityQuery{Limit: 2, Before: *next})
	if !slices.Equal(activityKinds(second), []Kind{ReportReady, SimFinished}) || next2 == nil {
		t.Fatalf("second %v", activityKinds(second))
	}
	last, next3, _ := st.Activity(ctx, uid, ActivityQuery{Limit: 2, Before: *next2})
	if !slices.Equal(activityKinds(last), []Kind{RaidPosted}) || next3 != nil {
		t.Fatalf("last %v next %v", activityKinds(last), next3)
	}
	// Exactly a page left is the last page.
	all, nextAll, _ := st.Activity(ctx, uid, ActivityQuery{Limit: 5})
	if len(all) != 5 || nextAll != nil {
		t.Fatalf("all %d next %v", len(all), nextAll)
	}
	ids := activityIDs(all)
	if !slices.IsSortedFunc(ids, func(x, y int64) int { return int(y - x) }) {
		t.Fatalf("not newest first %v", ids)
	}
}

func TestActivityFiltersByGroup(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "g")
	ctx := context.Background()
	recordAll(t, st, uid, RaidPosted, RaidShort, SimFinished, BandEntered, CoachReady, BillingFailed)
	cases := map[Group][]Kind{
		GroupRaids:   {RaidPosted},
		GroupOfficer: {RaidShort},
		GroupGear:    {BandEntered, SimFinished},
		GroupLogs:    {CoachReady},
		GroupAccount: {BillingFailed},
		GroupLadder:  {},
	}
	for g, want := range cases {
		rows, _, err := st.Activity(ctx, uid, ActivityQuery{Group: g, Limit: 30})
		if err != nil || !slices.Equal(activityKinds(rows), want) {
			t.Errorf("%s = %v %v, want %v", g, activityKinds(rows), err, want)
		}
	}
}

func TestUnreadCountAndMarkRead(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid, other := seedUser(t, pool, "u"), seedUser(t, pool, "v")
	ctx := context.Background()
	recordAll(t, st, uid, RaidPosted, SimFinished, ReportReady)
	recordAll(t, st, other, RaidPosted)
	if n, err := st.UnreadActivity(ctx, uid); err != nil || n != 3 {
		t.Fatalf("unread %d %v", n, err)
	}
	rows, _, _ := st.Activity(ctx, uid, ActivityQuery{Limit: 30})
	left, err := st.MarkActivityRead(ctx, uid, rows[1].ID)
	if err != nil || left != 1 {
		t.Fatalf("left %d %v", left, err)
	}
	rows, _, _ = st.Activity(ctx, uid, ActivityQuery{Limit: 30})
	if rows[0].Read || !rows[1].Read || !rows[2].Read {
		t.Fatalf("read %+v", rows)
	}
	if n, _ := st.UnreadActivity(ctx, other); n != 1 {
		t.Fatalf("another account changed: %d", n)
	}
	// Through the newest row reads the rest.
	if left, _ := st.MarkActivityRead(ctx, uid, rows[0].ID); left != 0 {
		t.Fatalf("left %d", left)
	}
}

func TestPruneActivityForgetsOldRows(t *testing.T) {
	pool := testPool(t)
	st := &Store{Pool: pool}
	uid := seedUser(t, pool, "r")
	ctx := context.Background()
	recordAll(t, st, uid, RaidPosted, SimFinished)
	exec(t, pool, `update user_activity set occurred_at = now() - interval '91 days' where kind = 'raid.posted'`)
	n, err := st.PruneActivity(ctx, time.Now(), 100)
	if err != nil || n != 1 {
		t.Fatalf("pruned %d %v", n, err)
	}
	rows, _, _ := st.Activity(ctx, uid, ActivityQuery{Limit: 30})
	if !slices.Equal(activityKinds(rows), []Kind{SimFinished}) {
		t.Fatalf("left %v", activityKinds(rows))
	}
}

func TestTheRouterRecordsOverTheStoreWithPushesOff(t *testing.T) {
	w := newGuildWorld(t)
	st := &Store{Pool: w.pool}
	ctx := context.Background()
	if err := st.SaveSettings(ctx, w.inMember, SettingsChange{Prefs: []PrefChange{
		{RaidStarting, ChannelPush, false}, {RaidStarting, ChannelDiscord, false},
	}}); err != nil {
		t.Fatal(err)
	}
	sink := &fakeSink{}
	r := &Router{Store: st, Sinks: map[Channel]Sink{ChannelPush: sink}, Inbox: st}
	e := Event{Kind: RaidStarting, Audience: Audience{Kind: ToRaidInLate, GuildID: w.guild, RaidID: w.raid},
		DedupeKey: fmt.Sprintf("raid.starting:%d", w.raid), Title: "MC starts in 30 minutes"}
	for range 2 {
		if err := r.Publish(ctx, e); err != nil {
			t.Fatal(err)
		}
	}
	if len(sink.got) != 1 || sink.got[0].uid != w.leader {
		t.Fatalf("sent %v", sink.got)
	}
	for _, uid := range []int64{w.leader, w.inMember} {
		if n, _ := st.UnreadActivity(ctx, uid); n != 1 {
			t.Fatalf("user %d unread %d", uid, n)
		}
	}
	if n, _ := st.UnreadActivity(ctx, w.out); n != 0 {
		t.Fatalf("out unread %d", n)
	}
}
